package com.kmaebashi.henacat.util;

import java.io.*;

public class FoundResponse extends Response {
	private String _location;

	public FoundResponse(String location) {
		this._location = location;
	}

	void writeLocationLine (OutputStream output) 
			throws IOException {
		Util.writeLine (output, "Location: " + _location);
	}

	@Override
	void writeStatusLine (OutputStream output) 
			throws IOException {
		Util.writeLine (output, "HTTP/1.1 302 Found");
	}

	@Override
	public void sendResponse (OutputStream output) 
			throws IOException {
		writeStatusLine (output);
		writeDateLine (output);
		writeServerLine (output);
		writeLocationLine (output);
		writeConnectionLine (output);
		writeNullLine (output);
	}

}

