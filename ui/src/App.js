import React, { useState, useEffect } from 'react';
import PlayersTableRank from './modules/PlayersTable';
import PlayerCard from './modules/PlayerCard'


function App(){
    return (
        <div>

            <PlayerCard position={"Delantero"} name={"Guido Enrique"} nickname={"Guidoti"} rank={4}/>
     {/*       <PlayersCardsList />*/}
        </div>
    )
}


export default App;
