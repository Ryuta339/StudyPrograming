package com.kmaebashi.henacat.util;

import java.io.*;

public class OkResponseHeader extends OkResponse {
	public OkResponseHeader (String ext) {
		super (ext);
	}

	@Override
	void writeContentTypeLine (OutputStream output, String contentType) 
			throws IOException {
		Util.writeLine (output, "Content-type: " + contentType);
	}

	@Override
	void writeContents (OutputStream output) 
			throws IOException {
		// write nothing
	}
}
