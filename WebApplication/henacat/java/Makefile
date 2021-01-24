JAVAC		= javac
JAVA		= java
SRC			= src
BIN			= bin
MAIN		= Main
PACKPATH	= com/kmaebashi/henacat


compile: $(SRC)/$(PACKPATH)/*/*.java
	$(JAVAC) -cp $(BIN) -d $(BIN) $^

exec:
	$(JAVA) -cp $(BIN) $(MAIN)


%.pack: $(SRC)/$(PACKPATH)/%/*.java $(SRC)/$(PACKPATH)/%/*/*.java
	$(JAVAC) -cp $(BIN) -d $(BIN) $^


clean:
	rm -rf $(BIN)/*
