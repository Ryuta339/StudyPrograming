import java.io.*;
import java.net.*;

public class TcpClient {
	public static void main (String[] args) throws Exception {
		try (Socket socket = new Socket("localhost", 8001);
				FileInputStream fis = new FileInputStream("client_send.txt");
				FileOutputStream fos = new FileOutputStream("client_recv.txt")) {
			
			int ch;
			// client_send.txt の内容をサーバーに送信
			OutputStream output = socket.getOutputStream();
			while ((ch = fis.read()) != -1) {
				output.write(ch);
			}
			// 終了を示すため、ゼロを送信
			output.write(0);

			// サーバーからの返信をclient_recv.txt に出力
			InputStream input = socket.getInputStream();
			while ((ch = input.read()) != 0) {
				fos.write(ch);
			}
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
}
