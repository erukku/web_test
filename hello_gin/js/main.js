
let app = new PIXI.Application({
    width:window.innerWidth,
    height:window.innerWidth*2/3,
    backgroundColor:0x2c3e50,
});

document.body.appendChild(app.view);


var graphics = new PIXI.Graphics();
graphics.beginFill(0x7fffd4, 0.5);

graphics.drawRect(0,window.innerWidth*2/3*(5/6),window.innerWidth,window.innerWidth*2/3*(1/6));
graphics.endFill();

graphics.beginFill(0xffffff, 0.5);
graphics.lineStyle(2,0xffffff)
graphics.drawRect(0,window.innerWidth*2/3*(5/6 + 1/12),window.innerWidth,window.innerWidth*2/3*(5/6 + 1/12)+30);
graphics.endFill();

graphics.beginFill(0xffff99, 0.5);
graphics.lineStyle(3, 0x777777);
graphics.drawRect(window.innerWidth/10,window.innerWidth*2/3*(5/6),window.innerWidth/10,window.innerWidth*2/3*(1/6));
graphics.endFill();

app.stage.addChild(graphics);

