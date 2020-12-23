JAVAC		= javac
JAVA		= java
SRC			= src
BIN			= bin
MAIN		= Main


compile: $(SRC)/*.java
	$(JAVAC) -d $(BIN) $^

exec:
	$(JAVA) -cp $(BIN) $(MAIN)
