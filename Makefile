BIN=target/1brc_go
SRC=src/main.go

${BIN}: ${SRC}
	go build -o $@ $<

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

.PHONY: run clean eval10k

eval10k: measurements_10K.txt ${BIN}
	@ln -fs $< measurements.txt
	@${BIN}
	@rm measurements.txt

eval100k: measurements_100K.txt ${BIN}
	@ln -fs $< measurements.txt
	@${BIN}
	@rm measurements.txt

eval1m: measurements_1M.txt ${BIN}
	@ln -fs $< measurements.txt
	@${BIN}
	@rm measurements.txt

eval10m: measurements_10M.txt ${BIN}
	@ln -fs $< measurements.txt
	@${BIN}
	@rm measurements.txt

run: ${BIN}
	${BIN}

clean:
	rm ${BIN}
	rm measurements*
