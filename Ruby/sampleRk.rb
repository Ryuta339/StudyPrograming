#! ruby -Ku
require "kconv"

class SampleRK
	def rk4 (time, data, max)
		epsilon = 0.001;
		data[0] = 0.0
		for i in 1..(max-2) do
			time[i] = time[i-1] + epsilon;
			k1 = func (time[i-1], data[i-1]);								# なぜかsyntax error
			k2 = func (time[i-1]+epsilon/2, data[i-1] + k1*epsilon/2);		# なぜかsyntax error
			k3 = func (time[i-1]+epsilon/2, data[i-1] + k2*epsilon/2);		# なぜかsyntax error
			k4 = func (time[i], data[i-1]+k3*epsilon);						# なぜかsyntax error
			data[i] = (k1 + 2*k2 + 2*k3 + k4) / 6;
		end
		return data;
	end
	
	def func (t, p)
		return 1 / (1 + p);
	end
end

rk = SampleRK.new;
MAX = 10000;
time = Array.new(MAX);
data = Array.new(MAX);
rk.rk4 (time, data, MAX);
for i in 0..(max-1) do
	puts time[i], data[i];
end
