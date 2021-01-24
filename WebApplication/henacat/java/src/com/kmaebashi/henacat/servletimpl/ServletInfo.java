package com.kmaebashi.henacat.servletimpl;

import com.kmaebashi.henacat.servlet.http.*;

public class ServletInfo {
	WebApplication _webApp;
	String _urlPattern;
	String _servletClassName;
	HttpServlet _servlet;

	public ServletInfo(WebApplication webApp, String urlPattern, 
			String servletClassName) {
		this._webApp = webApp;
		this._urlPattern = urlPattern;
		this._servletClassName = servletClassName;
	}

}
