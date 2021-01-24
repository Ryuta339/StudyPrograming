package com.kmaebashi.henacat.util;

import java.io.IOException;
import java.io.InputStream;
import java.io.OutputStream;
import java.text.DateFormat;
import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.HashMap;
import java.util.Locale;
import java.util.TimeZone;

public class Util {
	public static String readLine (InputStream input) throws IOException {
		int ch;
		String ret = "";
		while ((ch = input.read()) != -1) {
			if (ch == '\r') {
				// do nothing
			} else if (ch == '\n') {
				break;
			} else {
				ret += (char)ch;
			}
		}
		if (ch == -1)
			return null;
		return ret;
	}

	public static void writeLine (OutputStream output, String str) 
			throws IOException {
		for (char ch: str.toCharArray()) {
			output.write ((int)ch);
		}
		output.write('\r');
		output.write('\n');
	}

	public static String getDateStringUtc() {
		Calendar cal = Calendar.getInstance(TimeZone.getTimeZone("UTC"));
		DateFormat df = new SimpleDateFormat("EEE, dd MMM yyyy HH:mm:ss", Locale.US);
		df.setTimeZone (cal.getTimeZone());
		return df.format(cal.getTime()) + "GMT";
	}

	private static final HashMap <String, String> contentTypeMap =
		new HashMap<String, String> () {{
			put("html", "text/html");
			put("htm", "text/html");
			put("txt", "text/plain");
			put("css", "text/css");
			put("png", "image/png");
			put("jpg", "image/jpeg");
			put("jpeg", "image/jpeg");
			put("gif", "image/gif");
		}
	};

	public static String getContentType (String ext) {
		String ret = contentTypeMap.get (ext.toLowerCase());
		if (ret == null) {
			return "application/octet-stream";
		}
		return ret;
	}
}


