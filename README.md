# Toktik - Social Media Wall
A social media wall is an aggregated form of live social media feeds from multiple social media channels such as [**Tiktok**](https://www.tiktok.com), Twitter, Instagram, Facebook, Tumblr, Vimeo, YouTube, and many more. A social media wall is also known as a social media wall, **Tiktok wall**, Instagram wall, hashtag wall, content wall, and social feed wall.

Anyone can use social walls to display user-generated content or social media posts on websites and digital displays at a store, or during events, weddings, conferences, and so on.

This implementation would search for and stream video content dynamically from [Tiktok](https://www.tiktok.com) (without requiring a TikTok login or interaction). When a user enters a keyword, the social media wall performs a live search and streams relevant video results, allowing users to view a curated feed of videos without requiring API permissions or long-term data storage.

## Architecture
Toktik contains 2 parts:

| **Name**       | **Description**                      | **Language**                          |
|----------------|--------------------------------------|---------------------------------------|
| Frontend       | User interface for Toktik            | JavaScript + Vue.js                   |
| Backend        | Server-side logic and API for Toktik | Golang + Echo                         |

![Diagram](/public/diagram.png "mind map v1")

## Features
- Search videos based on keywords 
- Video stream (without requiring a TikTok login or interaction)
- User Interest Profiling (Tagging and Vectorization)
- Customization support (e.g., colors, layout)
- Web Embeddable

## Local Setup

Clone the repository and run on the root folder:
```bash
pnpm i
```

**Environment:**

```bash
cd server
cp .env.example .env
```

Request your RapidAPI key in https://rapidapi.com/tikwm-tikwm-default/api/tiktok-scraper7

**Setup server:**
```bash
go mod download
```

**Run development server:**
```
make dev
```

**For live reloading in dev mode:**
```
air
```

**Run web:**
```bash
cd ..
pnpm app:dev
```

## Snapshot

![Demo](/public/demo.png "demo v1")

## Contributing
We're really excited that you're interested in contributing to Toktik!