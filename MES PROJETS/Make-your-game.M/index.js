
const world = document.querySelector('#gameBoard');
const c = world.getContext('2d');

world.width = world.clientWidth;
world.height = world.clientHeight;

let frames=0;

const keys = {
    ArrowLeft: { pressed: false },
    ArrowRight: { pressed: false },
}

//Creation du joueur
class Player{
    constructor(){
        this.width=32; // largeur du player
        this.height=32; // hauteur du player
        this.velocity={
            x:0, // vitesse de déplacement axe des X
            y:0 // vitesse de déplacement axe des y
        }

        const image= new Image();
        image.src = 'space.png';
        image.src = './space.png';
        image.onload =()=>{
            this.image = image;
            this.width=48;
            this.height=48;
            this.position={
                x:world.width/2 - this.width/2,    // position du joueur par défaut centre
                y:world.height - this.height -10  // position du joueur par défaut centre
            }
        }
    }
    draw(){
        // le joueur sera un carré blanc
        /*c.fillStyle = 'white';
       c.fillRect(this.position.x,this.position.y,this.width,this.height);*/
       c.drawImage(this.image,
        this.position.x,
        this.position.y,
        this.width,
        this.height,
         );

    }
    // creer un missile
    shoot(){
        missiles.push(new Missile({
            position:{
                x:this.position.x + this.width/2,
                y : this.position.y
            },


        }))
    }


    update(){
        // A chaque mise à jour on dessine le joueur
        if(this.image) {
            if(keys.ArrowLeft.pressed && this.position.x >= 0 ){
                this.velocity.x = -5;
            }
            else if(keys.ArrowRight.pressed && this.position.x <= world.width - this.width){
                this.velocity.x = 5;
            }
            else{ this.velocity.x=0;
            }
            this.position.x += this.velocity.x;
            this.draw(); 

        }
       
    }

}

// creation des aliens
class Alien{
    constructor({position}){
        this.velocity={x:0, y:0 }
        const image= new Image();
        image.src = './ghost.png';
        image.onload =()=>{
            this.image = image;
            this.width=32;
            this.height=32  ;
            this.position= {
                x:position.x,
                y:position.y
            }
        }
        
    }
    // l'image qu'on veut utiliser creation du monstre 
    draw(){
        if(this.image){
        c.drawImage(this.image,this.position.x,this.position.y,this.width,this.height,);       
        }
    }

    update({velocity}){
        if(this.image){
        this.position.x += velocity.x;
        this.position.y += velocity.y;
        if(this.position.y + this.height >= world.height){
            console.log('You loose');
        }
        }
        this.draw();
    }
    shoot(alienMissiles){
        if(this.position){
            alienMissiles.push(new alienMissile({
                position:{
                    x:this.position.x,
                    y:this.position.y
                },
                velocity:{
                    x:0,
                    y:3
                }
            }))
        }
    }
}

//Création des misiiles pour le joueur
class Missile{
    constructor({position} ){
        this.position = position;
        this.velocity= {x:0, y:-5};
        this.width = 3;
        this.height = 10;
    }
    // construction du missile
    draw(){
        //c.save()
        c.fillStyle = 'red';
        c.fillRect(this.position.x, this.position.y, this.width, this.height);
        //c.restore()
    }
    update(){
       
        this.position.y += this.velocity.y;
        this.draw();
    }
}

// création missiles pour les aliens
class alienMissile{
    constructor({position,velocity}){
        this.position = position;
        this.velocity = velocity;
        this.width = 5;
        this.height =10;
    }
    draw(){
        
        c.fillStyle='yellow';
        c.fillRect(this.position.x,this.position.y,this.width,this.height)
       c.fill()
    
    }
    update(){
        this.position.x += this.velocity.x;
        this.position.y += this.velocity.y;
        this.draw()
    }
}

// Gestion des aliens à travers la classe grip c'est class permet de départager notre espace en jeu en ligne et en colonne
class Grid{
    constructor(){
        this.position={ x:0,y:0}
        this.velocity={x:1,y:0}
        this.invaders = [ ]
        let rows = Math.floor((world.height/34)*(1/3));
        const colums = Math.floor((world.width/34)*(2/3));
        this.height=rows*34;
        this.width = colums *34;
        for (let x=0;x<colums;x++){
			for(let y =0;y<rows;y++){
                this.invaders.push(new Alien({
                    position:{
                        x:x * 34,
                        y:y *34
                    }
                }))
            }
        }
    }
    update(){
        this.position.x += this.velocity.x;
        this.position.y += this.velocity.y;
        this.velocity.y =0;
        if(this.position.x + this.width  >= world.width || this.position.x == 0){
            this.velocity.x = -this.velocity.x;
            this.velocity.y = 34;
        }
        
        
    }
}
let alienMissiles = [];
const player = new Player();
const missiles = []; //l'ensemble des missiles présent sur notre tableau de jeu

let grids = [new Grid()];
let particules=[]; // tableau contenant toute les particules des aliens

//boucle d'animation

const animationLoop = () => {
    requestAnimationFrame(animationLoop);
    c.clearRect(0,0,world.width, world.height);
    player.update();
    // Pour faire le update de chaque missile
    missiles.forEach((missile, index) => {
        if(missile.position.y + missile.height <= 0) { // veriifie si le missible est dans l'espace de jeu
            setTimeout(() => {
                missiles.splice(index,1)} // on le supprime histoire de ne pas charger notre espace de jeu
                ,0)}
        else{missile.update();}
    })

    grids.forEach((grid) => {
      grid.update();

      // gerer les tires des aliens
      if(frames %50 ===0 && grid.invaders.length > 0) {
        grid.invaders[Math.floor(Math.random()*(grid.invaders.length))].shoot(alienMissiles);
        console.log(alienMissiles)
      }
        grid.invaders.forEach((invader, indexI) => {

            invader.update({velocity:grid.velocity});
            // pour tuer les aliens
            missiles.forEach((missile,indexM)=>{
                if(missile.position.y  <=  invader.position.y + invader.height &&
                   missile.position.y  >=  invader.position.y  &&
                   missile.position.x + missile.width >= invader.position.x &&
                   missile.position.x - missile.width <= invader.position.x + invader.width){
                /*    for(let i=0; i <12;i++){
                        particules.push(new Particule({
                            position:{
                                x:invader.position.x + invader.width/2,
                                y:invader.position.y + invader.height/2
                            },
                            velocity:{x:(Math.random()-0.5)*2,y:(Math.random()-0.5)*2},
                            radius:Math.random()*5+1,
                            color:'red'
                        }))
                    }*/
                setTimeout(()=>{
                    grid.invaders.splice(indexI,1);
                       
                    missiles.splice(indexM,1)
                   // if(grid.invaders.length === 0 && grids.length ==1 ){
                     //   grids.splice(indexGrid,1);
                    //    grids.push(new Grid());
                    //}
                },0)
                }
            })
        })
    })

    //Faire reference au lancé des missiles aliens
     alienMissiles.forEach((alienMissile,index) =>{
            if(alienMissile.position.y + alienMissile.height >=world.height){ 
                setTimeout(() =>{
                    alienMissiles.splice(index,1)} ,0);
                }
            else{alienMissile.update();}
        
           
        }) 


    frames++;
}

animationLoop();

//pour utilser le clavier en js on fait appel aux evenments

addEventListener('keydown', ({key}) =>{

    switch(key){
        case 'ArrowLeft':
            keys.ArrowLeft.pressed = true;
            console.log('gauche');
            break;
            case 'ArrowRight':
                keys.ArrowRight.pressed = true;
                console.log('droite');
                break;
    }
})


//Pour gérer le relachement de la touche

addEventListener('keyup', ({key}) => {
    switch(key) {
        case 'ArrowLeft':
            keys.ArrowLeft.pressed = false;
            console.log('gauche');
            break;
        case 'ArrowRight':
            keys.ArrowRight.pressed = false;
            console.log('droite');
            break;
            // permert d'afficher ce qui y'a dans notre tableau de missile lorsque le joueur click(gestion des missiles)
       case ' ':
            player.shoot();
            console.log(missiles);
            break;
    }
})

