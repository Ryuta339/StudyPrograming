package com.kmaebashi.henacat.util;

import java.io.*;

public class OkResponse extends Response {
	private InputStream _fis;
	private String _ext;


	public OkResponse(InputStream fis, String ext) {
		this._fis = fis;
		this._ext = ext;
	}

	protected OkResponse (String ext) {
		this._ext = ext;
	}

	void writeContents (OutputStream output) 
			throws IOException {
		int ch;
		while ((ch = _fis.read()) != -1) {
			output.write (ch);
		}
	}

	@Override
	void writeStatusLine (OutputStream output) 
			throws IOException {
		Util.writeLine (output, "HTTP/1.1 200 OK");
	}

	@Override
	public void sendResponse (OutputStream output)
			throws IOException {
		writeStatusLine (output);
		writeDateLine (output);
		writeServerLine (output);
		writeConnectionLine (output);
		writeContentTypeLine (output, _ext);
		writeNullLine (output);
		writeContents (output);
	}

}

