package com.kmaebashi.henacat.servletimpl;

import java.io.*;
import java.net.*;
import java.nio.file.*;
import java.util.*;

public class WebApplication {
	private static String WEBAPPS_DIR = "/usr/webapps";
	private static Map <String, WebApplication> webAppMap
		= new HashMap<String, WebApplication>();
	String _directory;
	ClassLoader _classLoader;
	private Map <String, ServletInfo> _servletMap
		= new HashMap <String, ServletInfo>();

	private WebApplication (String dir) throws MalformedURLException {
		this._directory = dir;
		FileSystem fs = FileSystems.getDefault();

		Path pathObj = fs.getPath(WEBAPPS_DIR + File.separator + dir);
		this._classLoader
			= URLClassLoader.newInstance(new URL[]{pathObj.toUri().toURL()});
	}

	public static WebApplication createInstance (String dir)
			throws MalformedURLException {
		WebApplication newApp = new WebApplication (dir);
		webAppMap.put (dir, newApp);

		return newApp;
	}

	public void addServlet (String urlPattern, String servletClassName) {
		_servletMap.put (urlPattern,
				new ServletInfo(this, urlPattern, servletClassName));
	}

	public ServletInfo searchServlet (String path) {
		return _servletMap.get(path);
	}

	public static WebApplication searchWebApplication (String dir) {
		return webAppMap .get(dir);
	}
}
