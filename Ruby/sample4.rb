#! ruby -Ku
require "kconv"

class HelloRuby
	def hello
		puts ("Hello, Ruby!");
	end
end
greeting = HelloRuby.new
greeting.hello
