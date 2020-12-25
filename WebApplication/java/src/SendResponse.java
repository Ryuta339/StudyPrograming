import java.io.*;

class SendResponse {
	static void sendOkResponse (OutputStream output, InputStream input, String ext) throws Exception {
		System.out.println ("ok");
		Util.writeLine (output, "HTTP/1.1 200 OK");
		Util.writeLine (output, "Date: " + Util.getDateStringUtc());
		Util.writeLine (output, "Server: Modoki/0.2");
		Util.writeLine (output, "Connection: close");
		Util.writeLine (output, "Content-type: " + Util.getContentType(ext));
		Util.writeLine (output, "");

		int ch;
		while ((ch = input.read()) != -1) {
			output.write(ch);
		}
		output.flush();
	}

	static void sendMovePermanentlyResponse (OutputStream output, String location) 
			throws Exception {
		System.out.println ("moved permanently");
		Util.writeLine (output, "HTTP/1.1 301 Moved Permanently");
		Util.writeLine (output, "Date: " + Util.getDateStringUtc());
		Util.writeLine (output, "Server: Modoki/0.2");
		Util.writeLine (output, "Location: " + location);
		Util.writeLine (output, "Connection: close");
		Util.writeLine (output, "");

	}

	static void sendNotFoundResponse (OutputStream output, String errorDocumentRoot)
			throws Exception {
		System.out.println ("not found");
		Util.writeLine (output, "HTTP/1.1 404 Not Found");
		Util.writeLine (output, "Date: " + Util.getDateStringUtc());
		Util.writeLine (output, "Server: Modoki/0.2");
		Util.writeLine (output, "Connection: close");
		Util.writeLine (output, "Content-type: text/html");
		Util.writeLine (output, "");

		try (InputStream input
				= new BufferedInputStream (new FileInputStream (errorDocumentRoot
						+ "/404.html"))) {
			int ch;
			System.out.println (errorDocumentRoot + "/404.html");
			while ((ch = input.read()) != -1) {
				output.write (ch);
			}
			output.flush();
		}
	}
}
