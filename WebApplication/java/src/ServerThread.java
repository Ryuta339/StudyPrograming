import java.io.*;
import java.net.*;
import java.text.*;
import java.util.*;

public class ServerThread implements Runnable {
	private static final String DOCUMENT_ROOT = "template";
	private Socket socket;

	private static final HashMap<String, String> contentTypeMap = 
		new HashMap<String, String>() {{
			put ("html", "text/html");
			put ("htm", "text/htm");
			put ("txt", "text/plain");
			put ("css", "text/css");
			put ("png", "image/png");
			put ("jpg", "image/jpg");
			put ("jpeg", "image/jpeg");
			put ("gif", "image/gif");
		}
	};
	
	private static String getContentType (String ext) {
		String ret = contentTypeMap.get (ext.toLowerCase());
		if (ret == null) {
			return "application/octet-stream";
		} else {
			return ret;
		}
	}

	private static String readLine (InputStream input) throws Exception {
		int ch;
		String ret = "";
		while ((ch = input.read()) != -1) {
			if (ch == '\r') {
				// do nothing
			} else if (ch == '\n') {
				break;
			} else {
				ret += (char) ch;
			}
		}
		if (ch == -1) {
			return null;
		} else {
			return ret;
		}
	}

	private static void writeLine (OutputStream output, String str) 
			throws Exception {
		for (char ch: str.toCharArray()) {
			output.write((int)ch);
		}
		output.write('\r');
		output.write('\n');
	}

	private static String getDateStringUtc() {
		Calendar cal = Calendar.getInstance(TimeZone.getTimeZone("UTC"));
		DateFormat df = new SimpleDateFormat("EEE, dd MMM yyyy HH:mm:ss", Locale.US);
		df.setTimeZone(cal.getTimeZone());
		return df.format(cal.getTime()) + " GMT";
	}

	@Override
	public void run () {
		OutputStream output;
		try {
			InputStream input = socket.getInputStream();

			String line;
			String path = null;
			String ext = null;
			while ((line = readLine(input)) != null) {
				if (line == "")
					break;
				if (line.startsWith("GET")) {
					path = line.split(" ")[1];
					String[] tmp = path.split("\\.");
					ext = tmp[tmp.length - 1];;;;
				}
			}
			output = socket.getOutputStream();

			// return response header
			writeLine (output, "HTTP/1.1 200 OK");
			writeLine (output, "Date: " + getDateStringUtc());
			writeLine (output, "Server: Modoki/0.1");
			writeLine (output, "Connection: close");
			writeLine (output, "Content-Type: " + getContentType (ext));
			writeLine (output, "");

			// return response body
			try (FileInputStream fis
					= new FileInputStream(DOCUMENT_ROOT + path); ) {
				int ch;
				while ((ch = fis.read()) != -1) {
					output.write(ch);
				}
			}
		} catch (Exception e) {
			e.printStackTrace ();
		} finally {
			try {
				socket.close();
			} catch (Exception e) {
				e.printStackTrace ();
			}
		}
	}

	ServerThread (Socket socket) {
		this.socket = socket;
	}
}
