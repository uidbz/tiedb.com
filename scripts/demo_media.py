#!/usr/bin/env python3
"""Build the demo library for tiedb.com screenshots.

The images are public-domain works from Wikimedia Commons. The albums are
fictional; their tracks are silent FLAC files with real tags (and realistic
durations) so the music player has something to show without shipping audio.

Requires: ffmpeg, ImageMagick (`magick`), and network access on the first run.
Set FORCE=1 to rebuild everything.
"""
import json
import os
import random
import shutil
import subprocess
import sys
import urllib.parse
import urllib.request

ROOT = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))
CACHE = os.path.join(ROOT, ".cache", "demo")
IMAGES = os.path.join(CACHE, "images")
MUSIC = os.path.join(CACHE, "music")
FORCE = os.environ.get("FORCE") == "1"

ART = [
    ("great-wave.jpg", "The_Great_Wave_off_Kanagawa.jpg", "The Great Wave off Kanagawa", "Katsushika Hokusai"),
    ("blue-marble.jpg", "The_Blue_Marble_(remastered).jpg", "The Blue Marble", "NASA / Apollo 17"),
    ("pillars-of-creation.jpg", "Pillars_of_Creation.jpg", "Pillars of Creation", "NASA / Hubble"),
    ("mona-lisa.jpg", "Mona_Lisa,_by_Leonardo_da_Vinci,_from_C2RMF_retouched.jpg", "Mona Lisa", "Leonardo da Vinci"),
    ("night-watch.jpg", "The_Night_Watch_-_HD.jpg", "The Night Watch", "Rembrandt"),
    ("girl-with-a-pearl-earring.jpg", "Girl_with_a_Pearl_Earring.jpg", "Girl with a Pearl Earring", "Johannes Vermeer"),
    ("tower-of-babel.jpg", "Pieter_Bruegel_the_Elder_-_The_Tower_of_Babel_(Vienna)_-_Google_Art_Project_-_edited.jpg", "The Tower of Babel", "Pieter Bruegel the Elder"),
    ("earthrise.jpg", "NASA-Apollo8-Dec24-Earthrise.jpg", "Earthrise", "NASA / Apollo 8"),
    ("apollo-17.jpg", "The_Earth_seen_from_Apollo_17.jpg", "The Blue Marble (Apollo 17)", "NASA / Apollo 17"),
    ("full-moon.jpg", "FullMoon2010.jpg", "Full Moon", "Gregory H. Revera"),
    ("deep-field.jpg", "Hubble_ultra_deep_field.jpg", "Hubble Ultra Deep Field", "NASA / ESA"),
    ("water-lilies.jpg", "Claude_Monet_-_Water_Lilies_-_1906,_Ryerson.jpg", "Water Lilies", "Claude Monet"),
    ("birth-of-venus.jpg", "Sandro_Botticelli_-_La_nascita_di_Venere_-_Google_Art_Project_-_edited.jpg", "The Birth of Venus", "Sandro Botticelli"),
    ("luncheon.jpg", "Pierre-Auguste_Renoir_-_Luncheon_of_the_Boating_Party_-_Google_Art_Project.jpg", "Luncheon of the Boating Party", "Pierre-Auguste Renoir"),
    ("the-scream.jpg", "Edvard_Munch_-_The_Scream_-_Google_Art_Project.jpg", "The Scream", "Edvard Munch"),
    ("wheatfield.jpg", "Vincent_van_Gogh_-_Wheatfield_with_crows_-_Google_Art_Project.jpg", "Wheatfield with Crows", "Vincent van Gogh"),
]

ALBUMS = [
    ("Slow Cooling", "The Annealers", 2026, "Ambient", ["First Frost", "Heat Haze", "Slow Cooling", "The Long Cool"], "pillars-of-creation.jpg", "#e6a23c"),
    ("Northern Relays", "Mira Kalden", 2024, "Electronic", ["Relay One", "Signal Hill", "Night Shift", "Northbound"], "full-moon.jpg", "#38bdf8"),
    ("Night Ferry", "Ostrander", 2019, "Jazz", ["Departure", "Harbor Lights", "The Crossing", "Last Ferry"], "night-watch.jpg", "#a78bfa"),
    ("Paper Lanterns", "Hana Sato Trio", 2021, "Jazz", ["Lanterns", "Rain on the Roof", "Two Streets", "Homecoming"], "great-wave.jpg", "#f59e0b"),
    ("Field Notes", "Elias Voss", 2022, "Folk", ["Meadow", "The Stone Wall", "Field Notes", "Evening"], "wheatfield.jpg", "#a3e635"),
    ("Low Tide", "The Bowline Quartet", 2018, "Classical", ["Low Tide", "Bowline", "Salt Air", "Slack Water"], "birth-of-venus.jpg", "#2dd4bf"),
    ("Static Gardens", "Neve Arden", 2025, "Ambient", ["Static Garden", "Glasshouse", "Slow Bloom", "After Rain"], "water-lilies.jpg", "#e879f9"),
    ("The Long Way Home", "Rook & Rowan", 2020, "Folk", ["The Long Way Home", "Paper Maps", "Riverbend", "Porch Light"], "wheatfield.jpg", "#fb7185"),
    ("Interference", "Kernel Panic", 2023, "Electronic", ["Carrier", "Interference", "Packet Loss", "Handshake"], "deep-field.jpg", "#ef4444"),
    ("Blue Hour", "Camille Roux", 2017, "Jazz", ["Blue Hour", "Rue des Fleurs", "Afterglow", "Metro"], "blue-marble.jpg", "#60a5fa"),
    ("Salt & Signal", "The Coastline", 2021, "Indie", ["Salt & Signal", "Radio Tower", "Low Cloud", "Tide Line"], "earthrise.jpg", "#34d399"),
    ("Afterglow", "Ines Marlow", 2024, "Pop", ["Afterglow", "Neon Rain", "Hold the Light", "The Last Song"], "birth-of-venus.jpg", "#f472b6"),
]


def run(cmd, **kw):
    subprocess.run(cmd, check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL, **kw)


def fetch_images():
    os.makedirs(IMAGES, exist_ok=True)
    credits = []
    for name, original, title, artist in ART:
        path = os.path.join(IMAGES, name)
        credits.append({"file": name, "title": title, "artist": artist, "source": "Wikimedia Commons (public domain)"})
        if os.path.exists(path) and not FORCE:
            continue
        url = "https://commons.wikimedia.org/wiki/Special:FilePath/" + urllib.parse.quote(original) + "?width=1400"
        req = urllib.request.Request(url, headers={"User-Agent": "tiedb.com demo corpus (screenshots)"})
        try:
            with urllib.request.urlopen(req, timeout=60) as r:
                data = r.read()
            if len(data) < 20000:
                raise RuntimeError("file too small")
            with open(path, "wb") as f:
                f.write(data)
            print("image:", name)
        except Exception as e:
            print("skip image:", name, e, file=sys.stderr)
    with open(os.path.join(IMAGES, "credits.json"), "w") as f:
        json.dump(credits, f, indent=2)


def make_music():
    random.seed(7)
    os.makedirs(MUSIC, exist_ok=True)
    title_font = "/usr/share/fonts/noto/NotoSans-Bold.ttf"
    body_font = "/usr/share/fonts/noto/NotoSans-Regular.ttf"
    for album, artist, year, genre, tracks, cover, hue in ALBUMS:
        d = os.path.join(MUSIC, artist, album)
        os.makedirs(d, exist_ok=True)
        cover_path = os.path.join(d, "cover.jpg")
        src = os.path.join(IMAGES, cover)
        if not os.path.exists(cover_path) or FORCE:
            run(["magick", src, "-resize", "900x900^", "-gravity", "center", "-extent", "900x900",
                 "-fill", "rgba(10,11,13,0.55)", "-draw", "rectangle 0,560 900,900",
                 "-font", title_font, "-pointsize", "54", "-fill", "white", "-annotate", "+54+670", album,
                 "-font", body_font, "-pointsize", "30", "-fill", "#d7d9de", "-annotate", "+54+720", artist,
                 "-fill", hue, "-draw", "rectangle 54,735 174,741", cover_path])
        for n, title in enumerate(tracks, 1):
            path = os.path.join(d, f"{n:02d} - {title}.flac")
            if os.path.exists(path) and not FORCE:
                continue
            dur = 150 + random.randint(0, 140)
            run(["ffmpeg", "-y", "-f", "lavfi", "-i", "anullsrc=r=44100:cl=stereo", "-t", str(dur),
                 "-c:a", "flac", "-compression_level", "8",
                 "-metadata", f"title={title}", "-metadata", f"artist={artist}",
                 "-metadata", f"album={album}", "-metadata", f"album_artist={artist}",
                 "-metadata", f"date={year}", "-metadata", f"track={n}/{len(tracks)}",
                 "-metadata", f"tracknumber={n}", "-metadata", f"genre={genre}", path])
        print("album:", artist, "/", album)


if __name__ == "__main__":
    for tool in ("ffmpeg", "magick"):
        if shutil.which(tool) is None:
            sys.exit(f"demo-media: {tool} is required")
    fetch_images()
    make_music()
    print("demo library:", CACHE)
