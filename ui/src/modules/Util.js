
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
    return require(`../assets/cartoon/${photoName}`)
}

function getTeamLogo(name){
    const logoName = name.toString().toLowerCase().replace(/\s/g,'') + ".png"
    return require(`../assets/teams/${logoName}`)
}


function getColor(position){
    // Set the color of the rectangle shape based on the value of the position prop
    // what a shit is this
    switch (position) {
        case 'delantero':
            return '#e90052'
        case 'defensor':
            return '#04f5ff'
        case 'volante':
            return '#00ff85'
        default:
            return '#38003c';
    }
}


export {getColor, getStars, getPhoto, getTeamLogo}