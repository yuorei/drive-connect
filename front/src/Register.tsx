import React, { useEffect, useState, createContext, useRef, useContext } from 'react';
import { Wrapper } from '@googlemaps/react-wrapper';
import useWindowSize from './components/useWindowSize';
import { Container } from '@mui/material';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Stepper from '@mui/material/Stepper';
import Step from '@mui/material/Step';
import StepLabel from '@mui/material/StepLabel';
import Button from '@mui/material/Button';
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import dayjs from 'dayjs';
import { useSearchParams } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';


function Register() {
  const [activeStep, setActiveStep] = useState(0);
  let deptPosition = { latitude: 0, longitude: 0 };
  let destPosition = { latitude: 0, longitude: 0 };
  const MAP_API_KEY = import.meta.env.VITE_MAP_API_KEY;

  const handleNext = () => {
    setActiveStep((prevActiveStep) => prevActiveStep + 1);
  };

  const steps = ['乗車地点の登録', '行き先の確認', '確認画面']

  const Step1: React.FC = () => {
    const MapContent: React.FC = () => {
      const ref = useRef<HTMLDivElement>(null);
      const map = useRef<google.maps.Map>();
      const [width, height] = useWindowSize();
      const mapWidth = width * 0.8;
      const mapHeight = height * 0.6;
      const [isAvailable, setAvailable] = useState(false);
      const mapMarker = useRef<google.maps.Marker[]>([]);

      const [searchParams] = useSearchParams();
      const currentLatitude = Number(searchParams.get('current_lat'));
      const currentLongitude = Number(searchParams.get('current_lng'));

      const isFirstRef = useRef(true);
      useEffect(() => {
        isFirstRef.current = false;
        if ('geolocation' in navigator) {
          setAvailable(true);
        }
      }, [isAvailable]);

      // 最初にMapを表示する時の設定
      const DEFAULT = {
        CENTER: {
          lat: 37.5206361,
          lng: 139.942454
        } as google.maps.LatLngLiteral,
        ZOOM: 16
      } as const;

      useEffect(() => {
        if (ref.current && !map.current) {
          console.log("Map initializing...")
          const option = {
            center: DEFAULT.CENTER,
            zoom: DEFAULT.ZOOM,
            disableDefaultUI: true,
            gestureHandling: "greedy",
          };
          const mapObj = new window.google.maps.Map(ref.current, option);
          map.current = mapObj;
          console.log("Map initialized!")
        }
      });

      useEffect(() => {
        if (!map) return;
        map.current?.addListener('click', (event: google.maps.MapMouseEvent) => {

          if (event.latLng === null) return
          const lat = event.latLng?.lat();
          const lng = event.latLng?.lng();
          console.log(lat, lng);
          const marker = new google.maps.Marker({
            position: { lat, lng },
          });
          console.log(marker);
          mapMarker.current[0]?.setMap(null);
          marker.setMap(map.current!);
          mapMarker.current[0] = marker;
          deptPosition = { latitude: lat, longitude: lng };
        });
      }, [map]);

      useEffect(() => {
        if (isAvailable && map) {
          map.current?.setCenter({
            lat: currentLatitude,
            lng: currentLongitude
          });
          const marker = new google.maps.Marker({
            position: {
              lat: currentLatitude,
              lng: currentLongitude
            },
          });
          mapMarker.current[0]?.setMap(null);
          marker.setMap(map.current!);
          mapMarker.current[0] = marker;
        }
      }, [map, currentLatitude, currentLongitude]);


      return (
        <div style={{ width: mapWidth, height: mapHeight }} ref={ref} />
      );
    };
    return (<>
      <h2>{steps[0]}</h2>
      <h3>どこから乗りますか？</h3>
      <Wrapper apiKey={MAP_API_KEY}>
        <MapContent />
      </Wrapper>
    </>)
  }

  const Step2: React.FC = () => {
    const MapContent: React.FC = () => {
      const ref = useRef<HTMLDivElement>(null);
      const [map, setMap] = useState({} as google.maps.Map);
      const [width, height] = useWindowSize();
      const mapWidth = width * 0.8;
      const mapHeight = height * 0.6;
      const [isAvailable, setAvailable] = useState(false);
      const mapMarker = useRef<google.maps.Marker[]>([]);

      const [searchParams] = useSearchParams();
      const destLatitude = Number(searchParams.get('dest_lat'));
      const destLongitude = Number(searchParams.get('dest_lng'));

      const isFirstRef = useRef(true);
      useEffect(() => {
        isFirstRef.current = false;
        if ('geolocation' in navigator) {
          setAvailable(true);
        }
      }, [isAvailable]);

      // 最初にMapを表示する時の設定
      const DEFAULT = {
        CENTER: {
          lat: 37.5206361,
          lng: 139.942454
        } as google.maps.LatLngLiteral,
        ZOOM: 16
      } as const;

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
          destPosition = { latitude: lat, longitude: lng };
        });
      }, [map]);

      useEffect(() => {
        if (isAvailable && map) {
          map.setCenter({
            lat: destLatitude,
            lng: destLongitude
          });
          const marker = new google.maps.Marker({
            position: {
              lat: destLatitude,
              lng: destLongitude
            },
          });
          mapMarker.current[0]?.setMap(null);
          marker.setMap(map);
          mapMarker.current[0] = marker;
        }
      }, [map, destLatitude, destLongitude]);


      return (
        <div style={{ width: mapWidth, height: mapHeight }} ref={ref} />
      );
    };
    return (
      <>
        <h2>{steps[1]}</h2>
        <h3>どこへ行きますか？</h3>
        <Wrapper apiKey={MAP_API_KEY}>
          <MapContent />
        </Wrapper>
      </>
    )
  }

  const [deptAddress, setDeptAddress] = useState('');
  const [destAddress, setDestAddress] = useState('');
  const [description, setDescription] = useState('');
  const [time, setTime] = useState('');

  // useEffect(async () => {
  //   const { Geocoder } = await google.maps.importLibrary("geocoding") as google.maps.GeocodingLibrary;
  //   Geocoder.geocode({
  //     location: {
  //       lat: deptPosition.latitude,
  //       lng: deptPosition.longitude
  //     }
  //   }, (results, status) => {
  //     if (status === 'OK') {
  //       setDeptAddress(results![0].formatted_address);
  //     }
  //   });
  //   Geocoder.geocode({
  //     location: {
  //       lat: destPosition.latitude,
  //       lng: destPosition.longitude
  //     }
  //   }, (results, status) => {
  //     if (status === 'OK') {
  //       setDestAddress(results![0].formatted_address);
  //     }
  //   });
  // }, [])
  
  const Step3: React.FC = () => {
    const navigate = useNavigate();
    const handleDescriptionChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      return setDescription(e.target.value);
    }
    const handleTimeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
      if(e === null) return;
      const date = dayjs(e.target.value).toISOString();
      return setTime(date);
    }
    const handleSubmit = () => {
      console.log(deptPosition);
      console.log(destPosition);
      console.log(description);
      console.log(time);
      return navigate('/');
     }
    return (
      <>
        <h2>{steps[2]}</h2>
        {/* <h3>確認画面</h3>
        <p>乗車地点</p>
        {{deptAddress}}
        <p>行き先</p>
        {{ destAddress }} */}
        <p>備考</p>
        <TextField
          id="outlined-basic"
          label="備考"
          variant="outlined"
          onChange={handleDescriptionChange}
        />
        <p>乗車希望時間</p>
        <TextField
          id="outlined-basic"
          label="乗車希望時間"
          variant="outlined"
          onChange={handleTimeChange}
        />
        <Button
          onClick={handleSubmit}
          sx={{ mt: 1, mb: 1 }}
        >
          登録する
        </Button>
      </>
    )
  }

  return (
    <>
      <Container>
        <Card
          sx={{
            position: 'fixed',
            top: 30,
            zIndex: 10,
            width: '90%',

          }}
        >
          <Stepper activeStep={activeStep}>
            {steps.map((label) => {
              const stepProps: { completed?: boolean } = {};
              const labelProps: { optional?: React.ReactNode } = {};
              return (
                <Step key={label} {...stepProps}>
                  <StepLabel {...labelProps}>{label}</StepLabel>
                </Step>
              );
            }
            )}
          </Stepper>
          <CardContent>
            <Box
              sx={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                textAlign: 'center',
              }}
            >
              {activeStep === 0 && <Step1 />}
              {activeStep === 1 && <Step2 />}
              {activeStep === 2 && <Step3 />}
              {activeStep !== 2 && <Button
                sx={{ mt: 1, mb: 1 }}
                onClick={handleNext}
              >
                次へ
              </Button>}
            </Box>
          </CardContent>
        </Card>
      </Container>

    </>
  )
}

export default Register
