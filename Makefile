main:
	git checkout main

branch:
	git checkout -b $(name)
	git push -u origin $(name)

push:
	git add .
	git commit -m "$(commit)"
	git push