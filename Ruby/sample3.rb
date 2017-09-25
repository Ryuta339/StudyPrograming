#! ruby -Ku
# 25th Sep 2017
# Ruby キモい
require "kconv"

array1 = Array["Red", "Red"];
print ('Array["Red, "Red"]' +"\n");
print array1[0].object_id, "\n";
print array1[1].object_id, "\n";

 array2 = Array.new 2, "Red";
 print ('Array.new (2, "Red")' + "\n");
 print array2[0].object_id, "\n";
 print array2[1].object_id, "\n";

array3 = Array.new(2){"Red"};
print ('Array.new(2){"Red"}' + "\n");
print array3[0].object_id, "\n";
print array3[1].object_id, "\n";

array4 = Array.new (array1);
print ('Array.new (array1)' + "\n");
print array4[0].object_id, "\n";
print array4[1].object_id, "\n";
