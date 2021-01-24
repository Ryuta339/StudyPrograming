package com.kmaebashi.henacat.webserver;

import com.kmaebashi.henacat.servletimpl.WebApplication;

import java.net.*;

public class Main {
	public static final String TESTBBS_TITLE = "testbbs";
	public static final int PORT = 8001;

	public static void main (String[] args) throws Exception {
		WebApplication app = WebApplication.createInstance(TESTBBS_TITLE);
		app.addServlet("/ShowBBS", "ShowBBS");
		app.addServlet("/PostBBS", "PostBBS");
		try (ServerSocket server = new ServerSocket(PORT)) {
			for (;;) {
				Socket socket = server.accept();
				System.out.println ("Connection accepted.");
				ServerThread serverThread = new ServerThread(socket);
				Thread thread = new Thread(serverThread);
				thread.start();
			}
		}
	}
}
