package com.kmaebashi.henacat.util;

import java.io.*;

public class NotFoundResponse extends Response {
	private String _errorDocumentRoot;

	public NotFoundResponse(String errorDocumentRoot) {
		this._errorDocumentRoot = errorDocumentRoot;
	}

	void writeErrorContents (OutputStream output) 
			throws IOException {
		try (InputStream fis
				= new BufferedInputStream (new FileInputStream
					(_errorDocumentRoot+"/404.html"))) {
			int ch;
			while ((ch = fis.read()) != -1) {
				output.write(ch);
			}
		}
	}

	@Override
	void writeStatusLine (OutputStream output) 
			throws IOException {
		Util.writeLine (output, "HTTP/1.1 404 Not Found");
	}

	@Override
	public void sendResponse (OutputStream output) 
			throws IOException {
		writeStatusLine (output);
		writeDateLine (output);
		writeServerLine (output);
		writeConnectionLine (output);
		writeContentTypeLine (output, "html");
		writeNullLine (output);
		writeErrorContents (output);
	}
}

