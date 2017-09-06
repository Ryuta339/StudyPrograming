#! ruby -Ku
require "kconv"

num = 0;

while num < 100 do
	print ("num = ")
	print (num)
	print (" ");
	case num % 5
		when 0 then
			print ("num mod 5 = 0\n");
		when 3 then
			print ("num mod 5 = 3\n");
		else
			print ("else\n");
		end
	num += 1;
end

