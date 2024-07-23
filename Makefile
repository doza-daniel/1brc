BIN=target/1brc_go
JAVA_BIN=target/CalculateAverage_thomaswue_image
SRC=src/main.go
JAVA_SRC=src/main/java/dev/morling/onebrc/CalculateAverage_thomaswue.java

${BIN}: ${SRC}
	go build -o $@ $<

${JAVA_BIN}: ${JAVA_SRC}
	./prepare_thomaswue.sh

measurements_10K.txt:
	./create_measurements3.sh 10000
	mv measurements3.txt $@

measurements_100K.txt:
	./create_measurements3.sh 100000
	mv measurements3.txt $@

measurements_1M.txt:
	./create_measurements3.sh 1000000
	mv measurements3.txt $@

measurements_10M.txt:
	./create_measurements3.sh 10000000
	mv measurements3.txt $@

.PHONY: eval10k
eval10k: measurements_10K.txt ${BIN}
	@ln -fs $< measurements.txt
	@${BIN}

.PHONY: eval100k
eval100k: measurements_100K.txt ${BIN}
	@ln -fs $< measurements.txt
	@${BIN}

.PHONY: eval1m
eval1m: measurements_1M.txt ${BIN}
	@ln -fs $< measurements.txt
	@${BIN}

.PHONY: eval10m
eval10m: measurements_10M.txt ${BIN}
	@ln -fs $< measurements.txt
	@${BIN}

# java

.PHONY: eval10k_java
eval10k_java: measurements_10K.txt ${JAVA_BIN}
	@ln -fs $< measurements.txt
	@${BIN}

.PHONY: eval100k_java
eval100k_java: measurements_100K.txt ${JAVA_BIN}
	@ln -fs $< measurements.txt
	@${JAVA_BIN}

.PHONY: eval1m_java
eval1m_java: measurements_1M.txt ${JAVA_BIN}
	@ln -fs $< measurements.txt
	@${JAVA_BIN}

.PHONY: eval10m_java
eval10m_java: measurements_10M.txt ${JAVA_BIN}
	@ln -fs $< measurements.txt
	@${JAVA_BIN}

.PHONY: run
run: ${BIN}
	@${BIN}

.PHONY: clean
clean:
	rm ${BIN}
