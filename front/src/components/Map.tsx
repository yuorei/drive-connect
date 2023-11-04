import { useEffect, useRef, useState, useContext } from 'react';
import { Wrapper } from '@googlemaps/react-wrapper';
import useWindowSize from './useWindowSize';
import { PositionContext } from '../Ride';
import { MapContext } from '../Ride';
import { useNavigate } from 'react-router-dom';

const MAP_API_KEY = import.meta.env.VITE_MAP_API_KEY;

// 最初にMapを表示する時の設定
const DEFAULT = {
  CENTER: {
    lat: 37.5206361,
    lng: 139.942454
  } as google.maps.LatLngLiteral,
  ZOOM: 14
} as const;

// 中身
const Content: React.FC = () => {
  const ref = useRef<HTMLDivElement>(null);
  const { map, setMap } = useContext(MapContext);
  const [width, height] = useWindowSize();
  const [isAvailable, setAvailable] = useState(false);
  const mapMarker = useRef<google.maps.Marker[]>([]);
  
  const { latitude, longitude } = useContext(PositionContext);

  const navigate = useNavigate();
  
  const isFirstRef = useRef(true);
  useEffect(() => {
    isFirstRef.current = false;
    if ('geolocation' in navigator) {
      setAvailable(true);
    }
  }, [isAvailable]);

  useEffect(() => {
    if (ref.current && Object.keys(map).length === 0) {
      console.log("Map initializing...")
      const option = {
        center: DEFAULT.CENTER,
        zoom: DEFAULT.ZOOM,
        disableDefaultUI: true,
        gestureHandling: "greedy",
      };
      const mapObj = new window.google.maps.Map(ref.current, option);
      setMap(mapObj);
      console.log("Map initialized!")
    }
  }, []);
  const infoWindowRef = useRef<google.maps.InfoWindow | null>(
    new google.maps.InfoWindow({
      content: `<div id="go-here">ここに行きたい！</div>`,
      maxWidth: 200
    })
  );

  useEffect(() => {
    if (Object.keys(map).length === 0) return;
    map.addListener('click', (event: google.maps.MapMouseEvent) => {

      if (event.latLng === null) return
      const lat = event.latLng?.lat();
      const lng = event.latLng?.lng();
      console.log(lat, lng);
      const marker = new google.maps.Marker({
        position: { lat, lng },
      });
      console.log(marker);
      mapMarker.current[0]?.setMap(null);
      marker.setMap(map);
      mapMarker.current[0] = marker;
      infoWindowRef.current?.open({
        anchor: marker,
        map,
      });
      infoWindowRef.current?.addListener('domready', () => {
        const button = document.getElementById('go-here');
        button?.addEventListener('click', async() => {
          console.log('clicked');
          mapMarker.current[0].setMap(null);
          // ここでRegisterに飛ばす
          return navigate(`/ride/register?dest_lat=${lat}&dest_lng=${lng}&current_lat=${latitude}&current_lng=${longitude}`);
        });
      });
    });
  }, [map]);

  useEffect(() => {
    if (isAvailable && map) {
      map.setCenter({
        lat: latitude,
        lng: longitude
      });
    }
  }, [map, latitude, longitude]);


  return (
    <div style={{ width, height }} ref={ref} />
  );
};


function Map() {
  return (
    <>
      <div style={{top:0}}>
        <Wrapper apiKey={MAP_API_KEY}>
          <Content />
        </Wrapper>
      </div>
    </>
  )
}

export default Map
