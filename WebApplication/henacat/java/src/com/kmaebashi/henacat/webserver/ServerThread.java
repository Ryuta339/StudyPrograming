package com.kmaebashi.henacat.webserver;

import com.kmaebashi.henacat.servlet.http.HttpRequestMethod;
import com.kmaebashi.henacat.servletimpl.*;
import com.kmaebashi.henacat.util.*;
import java.io.*;
import java.net.*;
import java.nio.file.*;
import java.util.*;

public class ServerThread implements Runnable {
	private static final String DOCUMENT_ROOT = System.getProperty("user.dir") + "/template";
	private static final String ERROR_DOCUMENT = "error_document";
	private Socket _socket;

	public ServerThread(Socket socket) {
		this._socket = socket;
	}

	private static void addRequestHeader (Map<String, String> requestHeader, String line) {
		int colonPos = line.indexOf(':');
		if (colonPos == -1)
			return;

		String headerName = line.substring(0, colonPos).toUpperCase();
		String headerValue = line.substring(colonPos+1).trim();
		requestHeader.put(headerName, headerValue);
	}

	@Override
	public void run () {
		OutputStream output = null;
		System.out.println("Run");
		try {
			InputStream input = _socket.getInputStream();

			String line;
			String requestLine = null;
			HttpRequestMethod method = null;
			Map<String, String> requestHeader = new HashMap<String, String>();
			while ((line = Util.readLine(input)) != null) {
				if (line == "") {
					break;
				}
				if (line.startsWith("GET")) {
					method = HttpRequestMethod.GET;
					requestLine = line;
				} else if (line.startsWith("POST")) {
					method = HttpRequestMethod.POST;
					requestLine = line;
				} else {
					addRequestHeader (requestHeader, line);
				}
			}
			if (requestLine == null)
				return;

			String reqUri = MyURLDecoder.decode (requestLine.split(" ")[1],
					"UTF-8");
			String[] pathAndQuery = reqUri.split("\\?");
			String path = pathAndQuery[0];
			String query = null;
			if (pathAndQuery.length > 1) {
				query = pathAndQuery[1];
			}
			output = new BufferedOutputStream(_socket.getOutputStream());

			String appDir = path.substring(1).split("/")[0];
			System.out.println("Application directory is " + appDir);
			WebApplication webApp = WebApplication.searchWebApplication(appDir);
			if (webApp != null) {
				ServletInfo servletInfo
					= webApp.searchServlet (path.substring(appDir.length()+1));
				if (servletInfo != null) {
					ServletService.doService(method, query, servletInfo, 
							requestHeader, input, output);
					return;
				}
			}
			String ext = null;
			String[] tmp = reqUri.split("\\.");
			System.out.println("Request URI is " + reqUri);
			ext = tmp[tmp.length-1];

			if (path.endsWith("/")) {
				path += "index.html";
				ext = "html";
			}
			FileSystem fs = FileSystems.getDefault();
			Path pathObj = fs.getPath (DOCUMENT_ROOT + path);

			Path realPath;
			try {
				realPath = pathObj.toRealPath();
			} catch (NoSuchFileException e) {
				Response resp = new NotFoundResponse (ERROR_DOCUMENT);
				resp.sendResponse(output);
				return;
			}
			System.out.println("RealPath is " + realPath);
			if (!realPath.startsWith(DOCUMENT_ROOT)) {
				Response resp = new NotFoundResponse (ERROR_DOCUMENT);
				resp.sendResponse (output);
				return;
			} else if (Files.isDirectory(realPath)) {
				String host = requestHeader.get("HOST");
				String location = "http://"
					+ ((host != null) ? host : Constants.SERVER_NAME)
					+ path + "/";
				Response resp = new MoveResponse(location);
				resp.sendResponse (output);
				return;
			}
			try (InputStream fis
					= new BufferedInputStream(Files.newInputStream(realPath))) {
				Response resp = new OkResponse (fis, ext);
				resp.sendResponse(output);
			} catch (FileNotFoundException e) {
				Response resp = new NotFoundResponse (ERROR_DOCUMENT);
				resp.sendResponse(output);
			}
		} catch (Exception e) {
			e.printStackTrace ();
		} finally {
			try {
				if (output != null) {
					output.close();
				}
				_socket.close();
			} catch (Exception e) {
				e.printStackTrace ();
			}
		}
	}

}
