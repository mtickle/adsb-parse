

psql -h "localhost" -p "5432" -U "pi" -P "Boomer2025" -d "adsb" -c "\COPY (SELECT * FROM mv_flight) TO mv_flight.json WITH (FORMAT text, HEADER FALSE)"


locahost:5432:adsb:pi:Boomer2025