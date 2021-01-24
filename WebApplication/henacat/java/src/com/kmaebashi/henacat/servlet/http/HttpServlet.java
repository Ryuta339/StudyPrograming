package com.kmaebashi.henacat.servlet.http;

import com.kmaebashi.henacat.servlet.*;

import java.io.IOException;

public class HttpServlet {
	protected void doGet (HttpServletRequest req, HttpServletResponse resp) 
			throws ServletException, IOException {
	}

	protected void doPost (HttpServletRequest req, HttpServletResponse resp) 
			throws ServletException, IOException {
	}

	public void service (HttpServletRequest req, HttpServletResponse resp)
			throws ServletException, IOException {
		req.getMethod().service (this, req, resp);
	}
}

