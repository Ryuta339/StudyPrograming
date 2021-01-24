package com.kmaebashi.henacat.servlet.http;

import java.io.*;
import java.util.*;

public interface HttpServletRequest {
	HttpRequestMethod getMethod();
	String getParameter (String name);
	List<String> getParameterValues (String name);
	void setCharacterEncoding (String env) throws UnsupportedEncodingException;
}

