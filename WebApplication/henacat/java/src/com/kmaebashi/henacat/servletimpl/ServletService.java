package com.kmaebashi.henacat.servletimpl;

import java.io.*;
import java.util.*;
import com.kmaebashi.henacat.servlet.http.*;
import com.kmaebashi.henacat.util.*;

public class ServletService {
	private static HttpServlet createServlet (ServletInfo info) 
			throws ClassNotFoundException, ClassCastException,
			  IllegalAccessException, InstantiationException,
			 ExceptionInInitializerError, SecurityException {
		Class<?> clazz
			= info._webApp._classLoader.loadClass(info._servletClassName);
		return (HttpServlet)clazz.newInstance();
	}


	private static Map<String, List<String>> stringToMap (String str) {
		Map<String, List<String>> parameterMap
			= new HashMap<String, List<String>>();

		if (str == null)
			return parameterMap;

		// str is not null
		// パラメータは?key1=value1&key2=value2&... という形
		String[] paramArray = str.split("&");
		for (String param: paramArray) {
			String[] keyValue = param.split("=");
			if (parameterMap.containsKey(keyValue[0])) {
				// すでにkey がマップに登録済み
				parameterMap.get(keyValue[0]).add(keyValue[1]);
			} else {
				List<String> list = new ArrayList<String>();
				list.add(keyValue[1]);
				parameterMap.put(keyValue[0], list);
			}
		}
		return parameterMap;
	}

	private static String readToSize(InputStream input, int size) 
			throws IOException {
		int ch;
		StringBuilder sb = new StringBuilder();
		int readSize = 0;

		while (readSize < size && (ch = input.read()) != -1) {
			sb.append((char)ch);
			readSize ++;
		}
		return sb.toString();
	}

	public static void doService (
			HttpRequestMethod method, 
			String query,
			ServletInfo info,
			Map<String, String> requestHeader,
			InputStream input,
			OutputStream output) 
			throws Exception {
		if (info._servlet == null) {
			info._servlet = createServlet (info);
		}
		System.out.println("URL Pattern is " + info._urlPattern);
		System.out.println("Servlet Class is " + info._servletClassName);

		ByteArrayOutputStream outputBuffer = new ByteArrayOutputStream();
		HttpServletResponseImpl resp
			= new HttpServletResponseImpl (outputBuffer);

		HttpServletRequest req;
		switch (method) {
			case GET: {
					System.out.println ("Create GET request.");
					Map <String, List<String>> mp = stringToMap(query);
					req = new HttpServletRequestImpl (HttpRequestMethod.GET, mp);
				} break;
			case POST: {
					System.out.println ("Create POST request.");
					int contentLength
						= Integer.parseInt (requestHeader.get("CONTENT-LENGTH"));
					String line = readToSize(input, contentLength);
					Map<String, List<String>> mp = stringToMap(line);
					req = new HttpServletRequestImpl (HttpRequestMethod.POST, mp);
			   } break;
			default:
				throw new AssertionError ("BAD METHOD" + method.getMethodName());
		}

		info._servlet.service(req, resp);
		System.out.println (resp._contentType);

		switch (resp._status) {
			case HttpServletResponse.SC_OK:
				Response okresp = new OkResponseHeader (resp._contentType);
				okresp.sendResponse (output);
				try {
					resp._printWriter.flush();
				} catch (NullPointerException e) {
					resp.getWriter().flush();
				}
				byte[] outputBytes = outputBuffer.toByteArray();
				for (byte b: outputBytes) {
					output.write(b);
				}
				break;
			case HttpServletResponse.SC_FOUND:
				String redirectLocation;
				if (resp._redirectLocation.startsWith("/")) {
					String host = requestHeader.get("HOST");
					redirectLocation = "http://"
						+ ((host!=null) ? host: Constants.SERVER_NAME)
						+ resp._redirectLocation;
				} else {
					redirectLocation = resp._redirectLocation;
				}
				Response rdresp = new FoundResponse(redirectLocation);
				rdresp.sendResponse (output);
		}
	}
}

