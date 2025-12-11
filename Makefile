main:
	git checkout main
	clear

branch:
	git checkout -b $(name)
	git push -u origin $(name)
	clear

push:
	git add .
	git commit -m "solution"
	git push
	clear