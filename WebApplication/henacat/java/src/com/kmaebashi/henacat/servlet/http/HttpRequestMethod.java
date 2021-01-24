package com.kmaebashi.henacat.servlet.http;

import com.kmaebashi.henacat.servlet.*;

import java.io.IOException;

public enum HttpRequestMethod {
	GET("Get") {
		@Override
		public void service (HttpServlet servlet, HttpServletRequest req, HttpServletResponse resp)
				throws ServletException, IOException {
			servlet.doGet(req, resp);
		}
	},
	POST("Post") {
		@Override
		public void service (HttpServlet servlet, HttpServletRequest req, HttpServletResponse resp) 
				throws ServletException, IOException {
			servlet.doPost (req, resp);
		}
	};

	private final String _method;

	private HttpRequestMethod (String method) {
		this._method = method;
	}

	public String getMethodName () {
		return _method;
	}

	public abstract void service (HttpServlet servlet, HttpServletRequest req, HttpServletResponse resp)
			throws ServletException, IOException;
}

