#! ruby -Ku

print ("18..20 is\n");
for num in 18..20 do
	print ("num = "); print (num); print ( "\n");
end

print ("18...20 is\n");
for num in 18...20 do
	print ("num = "); print (num); print ("\n");
end

print ("\"Ax\"..\"Bc\" is \n");
for str in "Ax".."Bc" do
	print ("str = " + str + "\n");
end
