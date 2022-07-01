import React from "react";
import {Link} from "react-router-dom";
export default function Collection() {
    const [movies, setMovies] = React.useState([]);
    React.useEffect(() => { return async () =>{ 
        const url = "http://localhost:8000/movies/";
        const fetchMode = {
            method : 'GET',
            mode : 'cors'
        }
        const response = await fetch(url, fetchMode);
        setMovies(await response.json());
    } }, []);
    const moviesDiv = movies.map(item => {
        return (
            <Link to={`/play/${item.Id}`} >
            <div className="movie" id={item.Id} key={item.Id}>
            <p>{item.Name}</p>
            <p>{item.Director}</p>
            </div></Link>
        );
    });
    return (
        <div className="movies">
            {moviesDiv}
        </div>
    );
}