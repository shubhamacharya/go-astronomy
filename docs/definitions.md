- **Tropical Year** : A tropical year is the time it takes for the Earth to complete one orbit around the Sun, measured from one spring equinox to the next. It's about 365.24 days long and determines the length of our calendar year.

- **Julian Calendar**: An old calendar introduced by Julius Caesar in 45 BC, with a year lasting 365.25 days, including a leap year every 4 years.

- **Gregorian calendar**: The calendar system most of the world uses today. It was introduced by Pope Gregory XIII in 1582 to fix errors in the Julian calendar, making the year 365.2425 days long with leap years to keep our dates in sync with the seasons.

- **Civil Calendar**: The official calendar used by a country for everyday activities, often based on the Gregorian calendar today.

- **Fundamental epoch:** A chosen starting point in time for a particular system of measurement.
- **Greenwich mean noon:** Midday (12:00 PM) at the Greenwich meridian.
- **Greenwich meridian:** The prime meridian (0° longitude) that passes through Greenwich, UK, used as a reference for time zones.
- **Julian day number (or Julian date):** The number of days (and fraction of a day) that have elapsed since the fundamental epoch of Greenwich mean noon on January 1, 4713 BC.
- **UT (Universal Time):** A time standard based on the Earth's rotation, often used as a global reference time.
- **Civil day:** The common 24-hour day that typically begins at midnight in a local time zone.
- **Time zone 0:** The time zone centered on the Greenwich meridian, which uses UT.


Here's a **Markdown summary** with **clear definitions** and **relationships between time standards**, suitable for inclusion in Obsidian:

---

## 🌍 Time Standards and Conversions in Civil and Astronomical Use

### 🕰️ Universal Time (UT)

* **Basis**: Earth's rotation as observed at the Greenwich meridian (longitude 0°).
* **Usage**: Serves as a foundation for civil timekeeping.
* **Irregularity**: Earth’s rotation is not perfectly uniform.

---

### ⏱️ International Atomic Time (TAI)

* **Definition**: A time standard based on atomic clocks coordinated globally.
* **Source**: Managed by the Bureau International de l’Heure (Paris).
* **Accuracy**: Extremely uniform due to atomic standards.
* **Relation**:

  * TAI is **ahead** of UTC by a whole number of seconds.
  * Example (June 2010): `TAI − UTC = 34 seconds`.

---

### 🛰️ Coordinated Universal Time (UTC)

* **Definition**: Atomic time adjusted with **leap seconds** to stay within **0.9 seconds** of UT.
* **Leap Seconds**: Inserted occasionally (typically June or December) to match Earth's rotation.
* **Usage**:

  * Basis of **legal civil timekeeping** worldwide.
  * Broadcast by stations like:

    * DCF 77 (Germany)
    * MSF 60 (UK)
    * WWV (USA)
* **Relation**:

  * **UTC ≈ UT** (with ±0.9 seconds max deviation)
  * **UTC = TAI − (Leap Seconds)**

---

### 📡 GPS Time

* **Definition**: Atomic time used by the Global Positioning System.
* **Maintained by**: US Naval Observatory.
* **Difference**:

  * Started as **equal to UTC** on 1980-01-06.
  * Does **not** include leap seconds.
  * Example (June 2010): `GPS Time = UTC + 15 seconds`.

---

### ⏳ Greenwich Mean Time (GMT)

* **Legacy Term**: Often used interchangeably with UTC.
* **Caution**: Before 1925, GMT started at **midday**, so was **12 hours offset** from UT.
* **Modern Usage**: Treated as equivalent to UTC in casual and broadcast use (e.g., BBC World Service).

---

### 🌐 Terrestrial Time (TT)

* **Definition**: Uniform time scale used in precise astronomy.
* **Relation**:

  * `TT = TAI + 32.184 seconds`
* **Replaces**:

  * **Ephemeris Time (ET)** before 1984.
  * Known as **Terrestrial Dynamic Time (TDT)** until renamed in 1991.

---

### 🕓 Local Time and Time Zones

* **Standard Time**: Most countries define their **local civil time** based on **offsets from UT/UTC**.
* **Daylight Saving Time (DST)**:

  * Adds 1 hour to UT during summer months.
  * In the UK, called **British Summer Time (BST)**.
* **Reason**:

  * To align working hours with daylight.
* **Time Zones**:

  * Typically follow whole-hour offsets from UTC.
  * Adopted globally to maintain noon alignment with the Sun.

---

### 🔄 Summary of Time Standard Relationships

| Time Standard | Definition                | Relation                                 |
| ------------- | ------------------------- | ---------------------------------------- |
| **UT**        | Based on Earth rotation   | ≈ UTC                                    |
| **TAI**       | Atomic clocks             | UTC + leap seconds                       |
| **UTC**       | TAI with leap seconds     | UT ± 0.9 sec                             |
| **GPS Time**  | Atomic, no leap seconds   | UTC + X sec (X = number of leap seconds) |
| **TT**        | Uniform astronomical time | TAI + 32.184 sec                         |
| **GMT**       | Synonym for UTC (modern)  | Historically ≠ UT before 1925            |

---

Here's a concise and clear explanation of **timezone offsets** in Markdown format, suitable for use in Obsidian:

---

## 🕓 Time Zone Offset

### 📌 Definition

A **time zone offset** is the difference in **hours and minutes** between **local time** and **Coordinated Universal Time (UTC)**.

---

### 🧮 Format

Offsets are typically represented as:

```
UTC±[hh]:[mm]
```

* **Positive Offset (UTC+X)**: Time **ahead** of UTC (e.g., India, China, Japan).
* **Negative Offset (UTC−X)**: Time **behind** UTC (e.g., USA, Canada, Brazil).

---

### 🌍 Common Time Zone Offsets

| Region                  | Time Zone Abbreviation | Offset from UTC |
| ----------------------- | ---------------------- | --------------- |
| UK (Winter)             | GMT / UTC              | UTC+00:00       |
| UK (Summer)             | BST                    | UTC+01:00       |
| New York (Winter)       | EST                    | UTC−05:00       |
| New York (Summer)       | EDT                    | UTC−04:00       |
| India                   | IST                    | UTC+05:30       |
| China                   | CST (China)            | UTC+08:00       |
| Japan                   | JST                    | UTC+09:00       |
| Central Europe (Winter) | CET                    | UTC+01:00       |
| Central Europe (Summer) | CEST                   | UTC+02:00       |
| California (Winter)     | PST                    | UTC−08:00       |
| California (Summer)     | PDT                    | UTC−07:00       |

---

### 🧭 How It Works

**Local Time = UTC + Time Zone Offset**

Examples:

* `UTC = 12:00`

  * In India (UTC+05:30): `Local Time = 17:30`
  * In New York (UTC−05:00): `Local Time = 07:00`

---

### ☀️ Daylight Saving Time (DST)

* Some regions shift the clock **+1 hour during summer months**.
* Example: UK shifts from **GMT (UTC+0)** to **BST (UTC+1)** in summer.

---

### 🔁 Time Conversion Logic (Pseudocode)

```text
local_time = utc_time + timezone_offset
utc_time = local_time - timezone_offset
```

---

Here’s a **Markdown-formatted summary** of important terms, definitions, and relationships from your text on **Sidereal Time (ST)** and **Universal Time (UT)**:

---

# 🕰️ Sidereal Time and Universal Time

## 📘 Key Definitions

### **Universal Time (UT)**

* **Definition**: Based on the **apparent motion of the Sun** around the Earth.
* **Purpose**: Governs **local civil time** in any part of the world.
* **Day Definition**: 1 solar day = time between two successive **solar meridian transits** (Sun crossing observer’s meridian).

---

### **Sidereal Time (ST)**

* **Definition**: Time based on the **apparent motion of stars**. Formally, it is the **hour angle of the vernal equinox**.
* **Purpose**: Used by astronomers to track **stellar positions**.
* **Sidereal Clock**: A clock that ensures **stars return to the same position** after exactly 24 sidereal hours.
* **Nature**: Sidereal time advances **faster** than solar (universal) time.

---

## 🔁 Relationships

### 🧭 Solar vs. Sidereal Time

| Feature               | Solar Time (UT)                                   | Sidereal Time (ST)          |
| --------------------- | ------------------------------------------------- | --------------------------- |
| Reference Object      | Sun                                               | Stars                       |
| Basis of Day          | Sun-to-meridian crossing                          | Vernal equinox (star field) |
| Day Length            | \~24 hours                                        | \~23h 56m 4s                |
| Days per Year         | \~365.25 solar days                               | \~366.25 sidereal days      |
| Synchronization Point | Equal once a year (autumnal equinox, \~22 Sept)   |                             |
| Drift                 | Sidereal time gains \~4 minutes/day on solar time |                             |

---

### 🧮 Mathematical Relation (Conceptual)

* **1 sidereal day** = **(1 solar day) × (365.25 / 366.25)**
* Hence:
  `24 sidereal hours ≈ 23h 56m 4s of solar (UT) time`

---

## 📍 Key Events

* **Autumnal Equinox (\~22 Sept)**: ST and UT are equal.
* **6 months later**: ST leads UT by **12 hours**.
* **1 year later**: ST and UT realign.

---
