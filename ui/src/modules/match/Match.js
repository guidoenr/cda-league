import React, { useState, useEffect } from 'react';
import Team from './Team'
import PlayerSelector from './PlayerSelector'
import './Match.css'

const Match = () => {
    const [showTeams, setShowTeams] = useState(false);
    const [availablePlayers, setAvailablePlayers] = useState([]);
    const [Team1, setTeam1] = useState({});
    const [Team2, setTeam2] = useState({});

    const handleTeamGeneration = (selectedPlayers) => {
        const data = { players: selectedPlayers }
        fetch('http://localhost:8080/generateMatch/', {
            method: 'POST',
            body: JSON.stringify(data),
            headers: { 'Content-Type': 'application/json' }
        })
            .then(res => res.json())
            .then(response => {
                setTeam1(response.team1);
                setTeam2(response.team2);
                setShowTeams(true);
            })
            .catch(error => console.error('Error:', error));
    }

        return (
            <div>
                <PlayerSelector />
                <div className="match-container">

                    <Team className="team-container"
                        name="TEAM 1"
                        players={Team1.players}
                        totalPlayers={Team1.totalPlayers}
                        chanceOfWinning={Team1.chanceOfWinning}
                    />
                    <Team className="team-container"
                        name="TEAM 2"
                        players={Team2.players}
                        totalPlayers={Team2.totalPlayers}
                        chanceOfWinning={Team2.chanceOfWinning}
                    />
                </div>
            </div>
        );
}



export default Match;
