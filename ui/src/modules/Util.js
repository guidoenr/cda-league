
// searching how many stars will the player have
function getStars(rank){
    let stars = [];
    for (let i = 0; i < rank; i++) {
        stars.push("★");
    }
    return stars;
}

// searching how many stars will the player have
function getPhoto(nickname){
    const photoName = nickname.toString().toLowerCase() + ".jpeg"
    return require(`./assets/cartoon/${photoName}`)
}

function getColor(position){
    // Set the color of the rectangle shape based on the value of the position prop
    // what a shit is this
    switch (position) {
        case 'delantero':
            return 'red'
        case 'defensor':
            return 'blue'
        case 'volante':
            return 'green'
        default:
            return '#333';
    }
}

export {getColor, getStars, getPhoto}