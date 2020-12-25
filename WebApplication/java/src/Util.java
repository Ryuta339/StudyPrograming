import java.io.*;
import java.util.*;
import java.text.*;

class Util {
	static final HashMap<String, String> contentTypeMap = 
		new HashMap<String, String>() {{
			put ("html", "text/html");
			put ("htm", "text/htm");
			put ("txt", "text/plain");
			put ("css", "text/css");
			put ("png", "image/png");
			put ("jpg", "image/jpg");
			put ("jpeg", "image/jpeg");
			put ("gif", "image/gif");
		}
	};
	
	static String getContentType (String ext) {
		String ret = contentTypeMap.get (ext.toLowerCase());
		if (ret == null) {
			return "application/octet-stream";
		} else {
			return ret;
		}
	}

	static String readLine (InputStream input) throws Exception {
		int ch;
		String ret = "";
		while ((ch = input.read()) != -1) {
			if (ch == '\r') {
				// do nothing
			} else if (ch == '\n') {
				break;
			} else {
				ret += (char) ch;
			}
		}
		if (ch == -1) {
			return null;
		} else {
			return ret;
		}
	}

	static void writeLine (OutputStream output, String str) 
			throws Exception {
		for (char ch: str.toCharArray()) {
			output.write((int)ch);
		}
		output.write('\r');
		output.write('\n');
		output.flush();
	}

	static String getDateStringUtc() {
		Calendar cal = Calendar.getInstance(TimeZone.getTimeZone("UTC"));
		DateFormat df = new SimpleDateFormat("EEE, dd MMM yyyy HH:mm:ss", Locale.US);
		df.setTimeZone(cal.getTimeZone());
		return df.format(cal.getTime()) + " GMT";
	}


}
