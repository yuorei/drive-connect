import { useEffect, useState, createContext } from 'react';
import SearchBox from './components/SearchBox'
import Map from './components/Map';

export const PositionContext = createContext({ latitude: 0, longitude: 0 });
export const MapContext = createContext({} as {
  map: google.maps.Map;
  setMap: React.Dispatch<React.SetStateAction<google.maps.Map>>;
});

function Ride() {
  type Position = { latitude: number; longitude: number };
  const [position, setPosition] = useState({ latitude: 36.523671, longitude: 139.9391309 } as Position);
  const [ map, setMap ] = useState({} as google.maps.Map);

  useEffect(() => {
    console.log("Getting current position...")
    navigator.geolocation.getCurrentPosition(position => {
      console.log(position.coords);
      const { latitude, longitude } = position.coords;
      setPosition({ latitude, longitude });
    });
    console.log("Got current position!")
  }, []);

  return (
    <>
      <MapContext.Provider value={ { map, setMap } }>
      <PositionContext.Provider value={ position }>
        <SearchBox />
        <Map />
      </PositionContext.Provider>
      </MapContext.Provider>
    </>
  )
}

export default Ride
