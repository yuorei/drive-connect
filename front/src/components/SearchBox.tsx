import { useContext, useState } from 'react';
import { Paper, TextField, InputAdornment, List, ListItem, ListItemButton, ListItemText } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';
import { PositionContext, MapContext } from '../Ride';

const SearchBox = () => {
  const { latitude, longitude } = useContext(PositionContext);
  const { map } = useContext(MapContext);
  const [searchResult, setSearchResult] = useState<google.maps.places.PlaceResult[]>([])

  const handleSearch = async(event: React.ChangeEvent<HTMLInputElement>) => {
    console.log(event.target.value);

    const request = {
      location: new google.maps.LatLng(latitude, longitude),
      radius: 10000,
      keyword: event.target.value,
    };
    console.log(map)
    const { PlacesService } = await google.maps.importLibrary("places") as google.maps.PlacesLibrary;
    const service = new PlacesService(map);
    service.nearbySearch(request, (results, status) => {
      if (status === google.maps.places.PlacesServiceStatus.OK) {
        console.log(results);
        if(results === null) return setSearchResult([]);
        setSearchResult(results);
      }
    });
  }

  const jumpTo = (place: google.maps.places.PlaceResult) => {
    if(!place.geometry?.location) return;
    const lat = place.geometry?.location.lat() || 0;
    const lng = place.geometry?.location.lng() || 0;
    map.setCenter({ lat, lng });
    map.setZoom(16);
    setSearchResult([]);
  }

  return (
    <Paper
      sx={{
        position: 'fixed',
        top: 30,
        left: 10,
        right: 0,
        zIndex: 10,
        width: '90%',
      }}
    >
      <TextField
        id="standard-basic"
        fullWidth
        onChange={handleSearch}
        InputProps={{
          startAdornment: (
            <InputAdornment position="start">
              <SearchIcon />
            </InputAdornment>
          ),
        }}
      />
      {searchResult.length > 0 && searchResult.length < 10 &&
        <List>
          {searchResult.map((result) => (
            <ListItem disablePadding>
              <ListItemButton>
                <ListItemText
                  primary={result.name}
                  onClick={() => jumpTo(result)}
                />
              </ListItemButton>
            </ListItem>
          ))
          }
        </List>
      }
    </Paper>
  );
};

export default SearchBox;