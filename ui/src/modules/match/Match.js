import React, { useState, useEffect } from 'react';
import Team from './Team'
import './Match.css'

const Match = () => {
    const [team1, setTeam1] = useState({});
    const [team2, setTeam2] = useState({});

    useEffect(() => {
        const fetchPlayers = async () => {
            try {
                const response = await fetch('http://localhost:8080/match');
                const data = await response.json();
                setTeam1(data.team1)
                setTeam2(data.team2)
            } catch (error) {
                console.error(error);
            }
        };
        fetchPlayers();
    }, []);


    console.log(team1, team2)
    return (
        <div className="match-container">
            <Team {...team1} />
            <Team {...team2} />
        </div>
    );
}

export default Match;
