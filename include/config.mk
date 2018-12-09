
i2pd_dat?=i2pd_dat

str=!\`@\#$%^&*()-_+={}[]|/:;.,<>~\"

gen=apg -n 1 -E '$(str)' -m 3 -x 3

SAVE_README_LINES=17

name:
	@echo -n "attacker=" | tee attack
	$(gen) | tee -a attack
