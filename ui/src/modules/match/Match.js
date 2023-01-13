import React, { useState, useEffect } from 'react';
import PlayersList from '../player/PlayersList';
import './Match.css'

function MatchView() {
    const [team1Players, setTeam1Players] = useState([]);
    const [team2Players, setTeam2Players] = useState([]);

    useEffect(() => {
        // Aquí podrías hacer una llamada a tu API para obtener
        // los jugadores de cada equipo y actualizar el estado
        // setTeam1Players y setTeam2Players
    }, []);

    return (
        <div className="match-container">
            <div className="team-container">
                <h2>Team 1</h2>
                <PlayersList players={team1Players} />
            </div>
            <div className="team-container">
                <h2>Team 2</h2>
                <PlayersList players={team2Players} />
            </div>
        </div>
    );
}

export default MatchView;
