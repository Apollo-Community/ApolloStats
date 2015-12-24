ApolloStats
--------------------------------------------------------------------------------

Webpage for showing various stats from the Apollo Station SS13 game database.
With heavy inspiration from other servers' webpages such as [Goon](http://goonhub.com/) and [/vg/station](http://ss13.pomf.se/index.php/bans).

TODO
--------------------------------------------------------------------------------

Command line interface:
- Default options to run silent and serve pages on a standard port (not 80).
- Version flag to display current version.
- Update flag? Atempts to update the binary with a new released version from
  the main repo, by doing some magic with github.
- Investigate how the host would like to run this service.

Templates:
- Show some nice error pages.
- Use existing assets from the forum, such as css and logos.
- Replace links for tmp static assets with ext. assets to forum server.

Good looking lists:
- Replace tables with formatted lists.
- Format bans as: Player (ip, cid) got bantype until time by admin for reason.
- Format account items as: Player was awarded item, at time.
- Format deaths as: Name (job) died at time, in room (pos) with damage.
- Format round stats as: A round of gamemode (ended at time) has ended after duration.

Account items:
- Would be nice to show why a player got an item too.

Heat maps:
- Really nice if we could show a heatmap of deaths.
- Ask @HiddenKn how he made his python version.

Game map:
- Huge, zoomable map of the main station (only).
- Store the map as picture tiles?
- Investigate how goon made their map.
- Need to rebuild the map after any new map changes from a commit.
