package com.kmaebashi.henacat.servletimpl;

import com.kmaebashi.henacat.servlet.http.*;
import com.kmaebashi.henacat.util.*;
import java.io.*;
import java.nio.charset.*;
import java.util.*;

public class HttpServletRequestImpl implements HttpServletRequest {
	private HttpRequestMethod _method;
	private String _characterEncoding = "ISO-8859-1";
	private Map<String, List<String>> _parameterMap;

	HttpServletRequestImpl (HttpRequestMethod method, Map<String, List<String>> parameterMap) {
		this._method = method;
		this._parameterMap = parameterMap;
	}

	@Override
	public HttpRequestMethod getMethod() {
		return _method;
	}

	@Override
	public String getParameter (String name) {
		List<String> values = getParameterValues(name);
		if (values == null) {
			// No such element
			return null;
		}
		return values.get(0);
	}

	@Override
	public List<String> getParameterValues(String name) {
		List<String> values = _parameterMap.get(name);
		if (values == null) {
			// No such element
			return null;
		}
		List<String> decoded = new ArrayList<String>();
		try {
			for (String v: values) {
				decoded.add(MyURLDecoder.decode(v, _characterEncoding));
			}
		} catch (UnsupportedEncodingException e) {
			throw new AssertionError (e);
		}
		return decoded;
	}

	@Override
	public void setCharacterEncoding (String env) 
			throws UnsupportedEncodingException {
		if (!Charset.isSupported(env)) {
			throw new UnsupportedEncodingException("encoding.." + env);
		}
		this._characterEncoding = env;
	}

}

