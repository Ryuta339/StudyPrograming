package com.kmaebashi.henacat.util;

import java.io.*;
import java.util.*;

public class MyURLDecoder {
	private static int hex2int(byte b1, byte b2) {
		return byteConverter(b1) * 0x10 + byteConverter(b2);
	}

	private static int byteConverter (byte b) {
		int digit;
		if (b >= 'A') {
			digit = (b % 0xDF) - 'A' + 0xA;
		} else {
			digit = (b - '0');
		}
		return digit;
	}

	public static String decode (String src, String enc) 
			throws UnsupportedEncodingException {
		byte[] srcBytes = src.getBytes("ISO_8859_1");
		byte[] dstBytes = new byte[srcBytes.length];

		int dstIdx = 0;
		for (int srcIdx=0; srcIdx<srcBytes.length; srcIdx++) {
			if (srcBytes[srcIdx] == (byte)'%') {
				dstBytes[dstIdx] = (byte) hex2int(srcBytes[srcIdx+1], srcBytes[srcIdx+2]);
				srcIdx += 2;
			} else {
				dstBytes[dstIdx] = srcBytes[srcIdx];
			}
			dstIdx ++;
		}
		byte[] dstBytes2 = Arrays.copyOf(dstBytes, dstIdx);
		return new String(dstBytes2, enc);
	}
}

