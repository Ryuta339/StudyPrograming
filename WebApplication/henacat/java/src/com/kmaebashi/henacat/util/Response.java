package com.kmaebashi.henacat.util;

import java.io.*;

abstract public class Response {
	void writeDateLine (OutputStream output) 
			throws IOException {
		Util.writeLine (output, "Date: " + Util.getDateStringUtc());
	}
	void writeServerLine (OutputStream output) 
			throws IOException {
		Util.writeLine (output, "Server: Henacat/0.1");
	}
	void writeConnectionLine (OutputStream output) 
			throws IOException {
		Util.writeLine (output, "Connection: close");
	}
	void writeContentTypeLine (OutputStream output, String ext) 
			throws IOException {
		Util.writeLine (output, "Content-type: " + Util.getContentType(ext));
	}
	void writeNullLine (OutputStream output) 
			throws IOException {
		Util.writeLine (output, "");
	}

	abstract void writeStatusLine (OutputStream output) throws IOException;
	abstract public void sendResponse (OutputStream output) throws IOException;
}
