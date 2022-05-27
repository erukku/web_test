
let app = new PIXI.Application({
    width:window.innerWidth,
    height:window.innerWidth*2/3,
    backgroundColor:0x2c3e50,
});

document.body.appendChild(app.view);

document.addEventListener("keydown",keyevent)

app.stage.sortableChildren = true;
var graphics = new PIXI.Graphics();
graphics.beginFill(0x7fffd4, 0.5);

graphics.drawRect(0,window.innerWidth*2/3*(5/6),window.innerWidth,window.innerWidth*2/3*(1/6));
graphics.endFill();

graphics.beginFill(0xffffff, 0.5);
graphics.lineStyle(2,0xffffff,0.6)
graphics.drawRect(0,window.innerWidth*2/3*(5/6 + 1/12),window.innerWidth,1);
graphics.endFill();

app.stage.addChild(graphics);

var good = new PIXI.Graphics();

good.beginFill(0xffff99, 0.5);
good.lineStyle(3, 0x777777);
good.drawRect(window.innerWidth/10,window.innerWidth*2/3*(5/6),window.innerWidth/10,window.innerWidth*2/3*(1/6));
good.endFill();

good.zIndex = 100;

app.stage.addChild(good);



var sq = new PIXI.Graphics();
sq.beginFill(0xff0000);
sq.drawRect(0, 0, 50, 50);
sq.endFill();
sq.x = 100;
sq.y = 100;
app.stage.addChild(sq);
var speed = 1;
var keyflag = 0;
var nodelist = [];
app.load(startup())
function startup(){
    app.ticker.add(function(delta)
    {
        if(keyflag == 1){
            sq.x += 1;
        }
        else if(keyflag == 2){
            sq.y += 1;
        }
        else if(keyflag == 3){
            sq.x -= 1;
        }
        else if(keyflag == 4){
            sq.y -= 1;
        }
        
    })

    var count = 0;
    app.ticker.add(function(delta)
    {
        count += 1
        if (count % 60 == 0){
            generateNode()
        }
        moveNode()
        
    })
    
}

function generateNode(){
    var circle = new PIXI.Graphics();
    circle.beginFill(0xffffff);
    circle.drawCircle(0, 0, window.innerWidth*2/3*(1/6)*1/5);
    circle.x = window.innerWidth + 60
    circle.y = window.innerWidth*2/3*(5/6) * (1 + 1/24)
    circle.endFill();
    app.stage.addChild(circle);
    nodelist.push(circle);
}

function moveNode(){
    nodelist.forEach(node=> 
        node.x -= 3
    )
    if (nodelist.length != 0){
        var first = nodelist[0]; 
        console.log(first.x);
        if (first.x < -30){
            app.stage.removeChild(nodelist[0])
            nodelist.shift();
        }
    }
}

function deleteAllNode(){
    var size = nodelist.length;
    for(let i = 0;i<size;i++){
        var node = nodelist.shift()
        app.stage.removeChild(node)
    } 
}

function keyevent(event){
    switch(event.key){
        case "ArrowRight":
            keyflag = 1;
            break;
        case "ArrowDown":
            keyflag = 2;
            break;
        case "ArrowLeft":
            keyflag = 3;
            break;
        case "ArrowUp":
            keyflag = 4;
            break;
        default:
            keyflag = 0;
            break;
    }
}