package com.kmaebashi.henacat.servletimpl;

import java.io.*;
import com.kmaebashi.henacat.servlet.http.*;

public class HttpServletResponseImpl implements HttpServletResponse{
	String _contentType = "application/octet-stream";
	private String _characterEncoding = "ISO-8859-1";
	private OutputStream _outputStream;
	PrintWriter _printWriter;
	int _status;
	String _redirectLocation;

	HttpServletResponseImpl (OutputStream output) {
		this._outputStream = output;
		this._status = SC_OK;
	}

	@Override
	public void setContentType(String contentType) {
		_contentType = contentType;
		String[] tmp = contentType.split (" *; *");
		if (tmp.length > 1) {
			String[] keyValue = tmp[1].split("=");
			if (keyValue.length == 2 && keyValue[0].equals("charset")) {
				setCharacterEncoding(keyValue[1]);
			}
		}
	}
	
	@Override
	public void setCharacterEncoding (String charset) {
		_characterEncoding = charset;
	}

	@Override
	public PrintWriter getWriter () throws IOException {
		_printWriter
			= new PrintWriter(new OutputStreamWriter(_outputStream,
						_characterEncoding));
		return _printWriter;
	}

	@Override
	public void sendRedirect (String location) {
		_redirectLocation = location;
		setStatus(SC_FOUND);
	}

	@Override
	public void setStatus (int sc) {
		_status = sc;
	}
}

