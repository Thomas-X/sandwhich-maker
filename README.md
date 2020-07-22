# sandwhich-maker
Ever wanted to get a sandwhich with different breads, meats, tomatoes and lettuces? Now you can thanks to gRPC!

This project uses gRPC, Node, Golang to setup a set of microservices, all communicating and working in harmony together to deliver yummy sandwhiches to your CLI :).
I was able to make this project thanks to the courses I took (this is proof of what I learned).

## What does it do? 
Check the output, currently it only has data for Soleil bread with the following values:
```
2019/11/26 15:23:46 Created toppings client &{0xc0000c7500}
2019/11/26 15:23:46 Created toppings client &{0xc0000c7880}
2019/11/26 15:23:46 ButcherPort: 6050
2019/11/26 15:23:46 Got tres meats:<name:"roastbeef" price:1 > meats:<name:"salami" price:2 > lettuces:<name:"iceberg" price:9 > lettuces:<name:"kale" price:4 > tomatoes:<name:"dutch" price:2 > tomatoes:<name:"roma" price:11 > 
2019/11/26 15:23:46 Got bres bread:<name:"Soleil" amount:2 price:12 grainUsed:200 > 
```
## Setup
Make sure you have a *nix based terminal when running setup
```shell script
# In root of project
chmod +x ./start.sh
./start.sh
```
