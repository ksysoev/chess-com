package chesscom

// PlayerProfile contains details about a Chess.com player.
type PlayerProfile struct {
	// ID is the canonical API URL for this profile (always self-referencing).
	ID string `json:"@id"`
	// URL is the chess.com profile page URL.
	URL string `json:"url"`
	// Username is the player's username.
	Username string `json:"username"`
	// Name is the player's personal first and last name (optional).
	Name string `json:"name,omitempty"`
	// Title is the chess title abbreviation (e.g. "GM"), if any.
	Title string `json:"title,omitempty"`
	// Status is the account status: closed, closed:fair_play_violations, basic, premium, mod, staff.
	Status string `json:"status"`
	// Avatar is the URL of the player's 200x200 avatar image (optional).
	Avatar string `json:"avatar,omitempty"`
	// Location is the player's city or location (optional).
	Location string `json:"location,omitempty"`
	// Country is the API URL of the player's country profile.
	Country string `json:"country"`
	// TwitchURL is the player's Twitch.tv URL, if they are a streamer.
	TwitchURL string `json:"twitch_url,omitempty"`
	// PlayerID is the non-changing Chess.com ID of this player.
	PlayerID int `json:"player_id"`
	// Joined is the Unix timestamp of registration on Chess.com.
	Joined int64 `json:"joined"`
	// LastOnline is the Unix timestamp of the most recent login.
	LastOnline int64 `json:"last_online"`
	// Followers is the number of players tracking this player's activity.
	Followers int `json:"followers"`
	// FIDE is the player's FIDE rating.
	FIDE int `json:"fide,omitempty"`
	// IsStreamer indicates whether the player is a Chess.com streamer.
	IsStreamer bool `json:"is_streamer"`
}

// TitledPlayers is the response for the titled players endpoint.
type TitledPlayers struct {
	// Players is the array of usernames for players with the requested title.
	Players []string `json:"players"`
}

// RatingStats holds the last/best rating and record for a specific game type.
type RatingStats struct {
	Last       *RatingPoint     `json:"last,omitempty"`
	Best       *BestRating      `json:"best,omitempty"`
	Record     *GameRecord      `json:"record,omitempty"`
	Tournament *TournamentStats `json:"tournament,omitempty"`
}

// RatingPoint is the most recent rating data.
type RatingPoint struct {
	// Date is the Unix timestamp of the last rated game finished.
	Date int64 `json:"date"`
	// Rating is the most-recent rating.
	Rating int `json:"rating"`
	// RD is the Glicko RD value.
	RD int `json:"rd"`
}

// BestRating is the best rating achieved by a win.
type BestRating struct {
	Game   string `json:"game"`
	Date   int64  `json:"date"`
	Rating int    `json:"rating"`
}

// GameRecord summarises all games played.
type GameRecord struct {
	// Win is the number of games won.
	Win int `json:"win"`
	// Loss is the number of games lost.
	Loss int `json:"loss"`
	// Draw is the number of games drawn.
	Draw int `json:"draw"`
	// TimePerMove is the integer number of seconds per average move.
	TimePerMove int `json:"time_per_move"`
	// TimeoutPercent is the timeout percentage in the last 90 days.
	TimeoutPercent float64 `json:"timeout_percent"`
}

// TournamentStats summarises tournament participation.
type TournamentStats struct {
	// Count is the number of tournaments joined.
	Count int `json:"count"`
	// Withdraw is the number of tournaments withdrawn from.
	Withdraw int `json:"withdraw"`
	// Points is the total number of points earned in tournaments.
	Points int `json:"points"`
	// HighestFinish is the best tournament place.
	HighestFinish int `json:"highest_finish"`
}

// TacticsStats holds highest and lowest tactics ratings.
type TacticsStats struct {
	Highest *RatingDate `json:"highest,omitempty"`
	Lowest  *RatingDate `json:"lowest,omitempty"`
}

// RatingDate pairs a rating value with a timestamp.
type RatingDate struct {
	// Rating is the rating value.
	Rating int `json:"rating"`
	// Date is the Unix timestamp.
	Date int64 `json:"date"`
}

// PuzzleRushStats holds puzzle rush scores.
type PuzzleRushStats struct {
	Daily *PuzzleRushScore `json:"daily,omitempty"`
	Best  *PuzzleRushScore `json:"best,omitempty"`
}

// PuzzleRushScore holds total attempts and score for a puzzle rush session.
type PuzzleRushScore struct {
	// TotalAttempts is the number of attempts.
	TotalAttempts int `json:"total_attempts"`
	// Score is the score achieved.
	Score int `json:"score"`
}

// PlayerStats contains ratings, win/loss, and other stats for a player.
type PlayerStats struct {
	ChessDaily    *RatingStats     `json:"chess_daily,omitempty"`
	Chess960Daily *RatingStats     `json:"chess960_daily,omitempty"`
	ChessRapid    *RatingStats     `json:"chess_rapid,omitempty"`
	ChessBlitz    *RatingStats     `json:"chess_blitz,omitempty"`
	ChessBullet   *RatingStats     `json:"chess_bullet,omitempty"`
	Tactics       *TacticsStats    `json:"tactics,omitempty"`
	Lessons       *TacticsStats    `json:"lessons,omitempty"`
	PuzzleRush    *PuzzleRushStats `json:"puzzle_rush,omitempty"`
}

// PlayerOnlineStatus is the response for the player online status endpoint.
type PlayerOnlineStatus struct {
	// Online is true if the player has been online in the last five minutes.
	Online bool `json:"online"`
}

// GamePlayer holds the details of one player in a completed game.
type GamePlayer struct {
	// ID is the API URL of this player's profile.
	ID string `json:"@id"`
	// Username is the player's username.
	Username string `json:"username"`
	// UUID is the member unique ID (live games only).
	UUID string `json:"uuid,omitempty"`
	// Result is the game result code for this player.
	Result string `json:"result"`
	// Rating is the player's rating after the game finished.
	Rating int `json:"rating"`
}

// GameAccuracies holds the accuracy scores for both players.
type GameAccuracies struct {
	// White is the white player's accuracy.
	White float64 `json:"white"`
	// Black is the black player's accuracy.
	Black float64 `json:"black"`
}

// Game represents a completed chess game.
type Game struct {
	Accuracies  *GameAccuracies `json:"accuracies,omitempty"`
	TimeControl string          `json:"time_control"`
	Rules       string          `json:"rules"`
	URL         string          `json:"url"`
	FEN         string          `json:"fen"`
	PGN         string          `json:"pgn"`
	Match       string          `json:"match,omitempty"`
	TimeClass   string          `json:"time_class"`
	Tournament  string          `json:"tournament,omitempty"`
	ECO         string          `json:"eco,omitempty"`
	Black       GamePlayer      `json:"black"`
	White       GamePlayer      `json:"white"`
	StartTime   int64           `json:"start_time,omitempty"`
	EndTime     int64           `json:"end_time,omitempty"`
	Rated       bool            `json:"rated,omitempty"`
}

// CurrentGame represents a currently-in-progress daily chess game.
type CurrentGame struct {
	// White is the URL of the white player's profile.
	White string `json:"white"`
	// Black is the URL of the black player's profile.
	Black string `json:"black"`
	// URL is the URL of this game.
	URL string `json:"url"`
	// FEN is the current FEN.
	FEN string `json:"fen"`
	// PGN is the current PGN.
	PGN string `json:"pgn"`
	// Turn indicates which player is to move: "white" or "black".
	Turn string `json:"turn"`
	// DrawOffer is the player who has made a draw offer (optional).
	DrawOffer string `json:"draw_offer,omitempty"`
	// TimeControl is the PGN-compliant time control.
	TimeControl string `json:"time_control"`
	// TimeClass is the time-per-move grouping.
	TimeClass string `json:"time_class"`
	// Rules is the game variant.
	Rules string `json:"rules"`
	// Tournament is the URL pointing to the tournament (if available).
	Tournament string `json:"tournament,omitempty"`
	// Match is the URL pointing to the team match (if available).
	Match string `json:"match,omitempty"`
	// MoveBy is the Unix timestamp of when the next move must be made.
	// Zero means the player-to-move is on vacation.
	MoveBy int64 `json:"move_by"`
	// LastActivity is the Unix timestamp of the last activity on the game.
	LastActivity int64 `json:"last_activity"`
	// StartTime is the Unix timestamp of game start.
	StartTime int64 `json:"start_time"`
}

// CurrentGames is the response from the current daily games endpoint.
type CurrentGames struct {
	// Games is the array of currently-in-progress daily games.
	Games []CurrentGame `json:"games"`
}

// ToMoveGame is an entry in the to-move daily chess games list.
type ToMoveGame struct {
	// URL is the URL of this game.
	URL string `json:"url"`
	// DrawOffer is true if this player has received a draw offer.
	DrawOffer bool `json:"draw_offer,omitempty"`
	// MoveBy is the Unix timestamp of when the move must be made by.
	// Zero means it is not this player's turn.
	MoveBy int64 `json:"move_by"`
	// LastActivity is the Unix timestamp of the last activity on the game.
	LastActivity int64 `json:"last_activity"`
}

// ToMoveGames is the response from the to-move endpoint.
type ToMoveGames struct {
	// Games is the array of daily games where it is the player's turn.
	Games []ToMoveGame `json:"games"`
}

// GameArchives is the response from the monthly archives list endpoint.
type GameArchives struct {
	// Archives is the array of URLs for monthly archives in ascending chronological order.
	Archives []string `json:"archives"`
}

// MonthlyGames is the response from a monthly archive endpoint.
type MonthlyGames struct {
	// Games is the array of completed games for the month.
	Games []Game `json:"games"`
}

// PlayerClubEntry holds one club in a player's list of clubs.
type PlayerClubEntry struct {
	// ID is the URL of the club's API endpoint.
	ID string `json:"@id"`
	// Name is the club's name.
	Name string `json:"name"`
	// URL is the club's website URL.
	URL string `json:"url"`
	// Icon is the club's icon URL.
	Icon string `json:"icon,omitempty"`
	// LastActivity is the Unix timestamp of last activity in the club.
	LastActivity int64 `json:"last_activity"`
	// Joined is the Unix timestamp of when the player joined the club.
	Joined int64 `json:"joined"`
}

// PlayerClubs is the response from the player clubs endpoint.
type PlayerClubs struct {
	// Clubs is the list of clubs the player is a member of.
	Clubs []PlayerClubEntry `json:"clubs"`
}

// MatchResult holds the results of a player's participation in a team match.
type MatchResult struct {
	// PlayedAsWhite is the result code for the game played as white.
	PlayedAsWhite string `json:"played_as_white,omitempty"`
	// PlayedAsBlack is the result code for the game played as black.
	PlayedAsBlack string `json:"played_as_black,omitempty"`
}

// PlayerMatchEntry is one team match entry in a player's match list.
type PlayerMatchEntry struct {
	Results *MatchResult `json:"results,omitempty"`
	Name    string       `json:"name"`
	URL     string       `json:"url"`
	ID      string       `json:"@id"`
	Club    string       `json:"club"`
	Board   string       `json:"board,omitempty"`
}

// PlayerMatches is the response from the player matches endpoint.
type PlayerMatches struct {
	// Finished is the list of completed matches.
	Finished []PlayerMatchEntry `json:"finished"`
	// InProgress is the list of ongoing matches.
	InProgress []PlayerMatchEntry `json:"in_progress"`
	// Registered is the list of upcoming matches the player is registered for.
	Registered []PlayerMatchEntry `json:"registered"`
}

// PlayerTournamentEntry is one tournament entry in a player's tournament list.
type PlayerTournamentEntry struct {
	// URL is the PubAPI URL of the tournament.
	URL string `json:"url"`
	// ID is the web URL of the tournament.
	ID string `json:"@id"`
	// Status is the final or current status of the player in the tournament.
	Status string `json:"status"`
	// Wins is the number of wins (finished only).
	Wins int `json:"wins,omitempty"`
	// Losses is the number of losses (finished only).
	Losses int `json:"losses,omitempty"`
	// Draws is the number of draws (finished only).
	Draws int `json:"draws,omitempty"`
	// PointsAwarded is the points awarded (finished only).
	PointsAwarded int `json:"points_awarded,omitempty"`
	// Placement is the finishing placement (finished only).
	Placement int `json:"placement,omitempty"`
	// TotalPlayers is the total number of players in the tournament (finished only).
	TotalPlayers int `json:"total_players,omitempty"`
}

// PlayerTournaments is the response from the player tournaments endpoint.
type PlayerTournaments struct {
	// Finished is the list of completed tournaments.
	Finished []PlayerTournamentEntry `json:"finished"`
	// InProgress is the list of ongoing tournaments.
	InProgress []PlayerTournamentEntry `json:"in_progress"`
	// Registered is the list of upcoming tournaments the player is registered for.
	Registered []PlayerTournamentEntry `json:"registered"`
}

// Club contains details about a Chess.com club.
type Club struct {
	// ID is the canonical API URL for this profile.
	ID string `json:"@id"`
	// Name is the human-readable name of this club.
	Name string `json:"name"`
	// Icon is the URL of the club's 200x200 icon image (optional).
	Icon string `json:"icon,omitempty"`
	// Country is the API URL of the club's country profile.
	Country string `json:"country"`
	// Visibility is whether the club is "public" or "private".
	Visibility string `json:"visibility"`
	// JoinRequest is the URL to submit a request to join this club.
	JoinRequest string `json:"join_request"`
	// Description is the text description of the club.
	Description string `json:"description,omitempty"`
	// Admin is the list of API URLs for the admin player profiles.
	Admin []string `json:"admin"`
	// ClubID is the non-changing Chess.com ID of this club.
	ClubID int `json:"club_id"`
	// AverageDailyRating is the average daily rating of members.
	AverageDailyRating int `json:"average_daily_rating"`
	// MembersCount is the total member count.
	MembersCount int `json:"members_count"`
	// Created is the Unix timestamp of club creation.
	Created int64 `json:"created"`
	// LastActivity is the Unix timestamp of the most recent post, match, etc.
	LastActivity int64 `json:"last_activity"`
}

// ClubMember holds a member's username and join timestamp.
type ClubMember struct {
	// Username is the member's username.
	Username string `json:"username"`
	// Joined is the Unix timestamp of when they joined.
	Joined int64 `json:"joined"`
}

// ClubMembers is the response from the club members endpoint.
type ClubMembers struct {
	// Weekly is the list of members active in the last week.
	Weekly []ClubMember `json:"weekly"`
	// Monthly is the list of members active in the last month.
	Monthly []ClubMember `json:"monthly"`
	// AllTime is the full list of all members.
	AllTime []ClubMember `json:"all_time"`
}

// ClubMatchEntry is one team match entry in a club's match list.
type ClubMatchEntry struct {
	// Name is the team match name.
	Name string `json:"name"`
	// ID is the URL pointing to the team match API endpoint.
	ID string `json:"@id"`
	// Opponent is the URL of the opposing club's API endpoint.
	Opponent string `json:"opponent"`
	// Result is the match result code for this club (finished only).
	Result string `json:"result,omitempty"`
	// TimeClass is the time control class: "daily" or a live variant.
	TimeClass string `json:"time_class"`
	// StartTime is the Unix timestamp of the match start (finished/in-progress only).
	StartTime int64 `json:"start_time,omitempty"`
}

// ClubMatches is the response from the club matches endpoint.
type ClubMatches struct {
	// Finished is the list of completed matches.
	Finished []ClubMatchEntry `json:"finished"`
	// InProgress is the list of ongoing matches.
	InProgress []ClubMatchEntry `json:"in_progress"`
	// Registered is the list of upcoming matches.
	Registered []ClubMatchEntry `json:"registered"`
}

// TournamentSettings holds the configuration settings for a tournament.
type TournamentSettings struct {
	// Type is the tournament type, e.g. "round_robin".
	Type string `json:"type"`
	// Rules is the game variant.
	Rules string `json:"rules"`
	// TimeClass is the time control class.
	TimeClass string `json:"time_class"`
	// TimeControl is the PGN-compliant time control string.
	TimeControl string `json:"time_control"`
	// InitialGroupSize is the initial group size.
	InitialGroupSize int `json:"initial_group_size"`
	// UserAdvanceCount is how many users advance per group.
	UserAdvanceCount int `json:"user_advance_count"`
	// WinnerPlaces is the number of winner places.
	WinnerPlaces int `json:"winner_places"`
	// RegisteredUserCount is the number of registered users.
	RegisteredUserCount int `json:"registered_user_count"`
	// GamesPerOpponent is the number of games per opponent.
	GamesPerOpponent int `json:"games_per_opponent"`
	// TotalRounds is the total number of rounds.
	TotalRounds int `json:"total_rounds"`
	// ConcurrentGamesPerOpponent is the number of concurrent games per opponent.
	ConcurrentGamesPerOpponent int `json:"concurrent_games_per_opponent"`
	// IsRated indicates whether the tournament is rated.
	IsRated bool `json:"is_rated"`
	// IsOfficial indicates whether the tournament is official.
	IsOfficial bool `json:"is_official"`
	// IsInviteOnly indicates whether the tournament is invite-only.
	IsInviteOnly bool `json:"is_invite_only"`
	// UseTiebreak indicates whether a tiebreak system is used.
	UseTiebreak bool `json:"use_tiebreak"`
	// AllowVacation indicates whether vacation is allowed.
	AllowVacation bool `json:"allow_vacation"`
}

// TournamentPlayer is a player entry in a tournament.
type TournamentPlayer struct {
	// Username is the player's username.
	Username string `json:"username"`
	// Status is the player's status in the tournament.
	Status string `json:"status"`
}

// Tournament contains details about a Chess.com tournament.
type Tournament struct {
	Name        string             `json:"name"`
	URL         string             `json:"url"`
	Description string             `json:"description"`
	Creator     string             `json:"creator"`
	Status      string             `json:"status"`
	Players     []TournamentPlayer `json:"players"`
	Rounds      []string           `json:"rounds"`
	Settings    TournamentSettings `json:"settings"`
	FinishTime  int64              `json:"finish_time,omitempty"`
}

// TournamentRoundPlayer is a player entry in a tournament round.
type TournamentRoundPlayer struct {
	// Username is the player's username.
	Username string `json:"username"`
	// IsAdvancing indicates whether the player is advancing to the next round.
	IsAdvancing bool `json:"is_advancing"`
}

// TournamentRound contains details about a single tournament round.
type TournamentRound struct {
	// Groups is the list of URLs for each group in this round.
	Groups []string `json:"groups"`
	// Players is the list of players in this round.
	Players []TournamentRoundPlayer `json:"players"`
}

// TournamentGroupPlayer is a player entry in a tournament round group.
type TournamentGroupPlayer struct {
	// Username is the player's username.
	Username string `json:"username"`
	// Points is the points earned in this group.
	Points float64 `json:"points"`
	// TieBreak is the tie-break points earned in this group.
	TieBreak float64 `json:"tie_break"`
	// IsAdvancing indicates whether the player is advancing.
	IsAdvancing bool `json:"is_advancing"`
}

// TournamentRoundGroup contains details about a tournament round group.
type TournamentRoundGroup struct {
	// FairPlayRemovals is the list of usernames removed for fair play violations.
	FairPlayRemovals []string `json:"fair_play_removals"`
	// Games is the list of games in this group.
	Games []CurrentGame `json:"games"`
	// Players is the list of players in this group with their standings.
	Players []TournamentGroupPlayer `json:"players"`
}

// TeamMatchSettings holds the configuration for a team match.
type TeamMatchSettings struct {
	// TimeClass is the time control class.
	TimeClass string `json:"time_class"`
	// TimeControl is the PGN-compliant time control string.
	TimeControl string `json:"time_control"`
	// InitialSetup is the initial board setup (for Chess960 etc.).
	InitialSetup string `json:"initial_setup,omitempty"`
	// Rules is the game variant.
	Rules string `json:"rules"`
	// MinTeamPlayers is the minimum number of players per team.
	MinTeamPlayers int `json:"min_team_players"`
	// MaxTeamPlayers is the maximum number of players per team.
	MaxTeamPlayers int `json:"max_team_players"`
	// MinRequiredGames is the minimum number of required games.
	MinRequiredGames int `json:"min_required_games"`
	// MinRating is the minimum player rating (registration phase only).
	MinRating int `json:"min_rating,omitempty"`
	// MaxRating is the maximum player rating (registration phase only).
	MaxRating int `json:"max_rating,omitempty"`
	// Autostart indicates whether the match starts automatically.
	Autostart bool `json:"autostart"`
}

// TeamMatchPlayer is a player entry in a team match.
type TeamMatchPlayer struct {
	// Username is the player's username.
	Username string `json:"username"`
	// Board is the URL of the player's board.
	Board string `json:"board"`
	// Stats is the URL to the player's stats (finished matches).
	Stats string `json:"stats,omitempty"`
	// Status is the account status (registration phase).
	Status string `json:"status,omitempty"`
	// PlayedAsWhite is the result code for the game played as white.
	PlayedAsWhite string `json:"played_as_white,omitempty"`
	// PlayedAsBlack is the result code for the game played as black.
	PlayedAsBlack string `json:"played_as_black,omitempty"`
	// Rating is the player's rating (registration phase).
	Rating int `json:"rating,omitempty"`
	// RD is the Glicko RD value (registration phase).
	RD float64 `json:"rd,omitempty"`
	// TimeoutPercent is the player's timeout percentage.
	TimeoutPercent float64 `json:"timeout_percent,omitempty"`
}

// TeamMatchTeam holds the details for one side of a team match.
type TeamMatchTeam struct {
	// ID is the API URL of the club profile.
	ID string `json:"@id"`
	// URL is the web URL of the club profile.
	URL string `json:"url,omitempty"`
	// Name is the club name.
	Name string `json:"name"`
	// Result is the match result for this team (finished only).
	Result string `json:"result,omitempty"`
	// FairPlayRemovals is the list of usernames removed for fair play violations.
	FairPlayRemovals []string `json:"fair_play_removals,omitempty"`
	// Players is the list of players on this team.
	Players []TeamMatchPlayer `json:"players"`
	// Score is the team's score.
	Score float64 `json:"score"`
}

// TeamMatch contains details about a daily team match.
type TeamMatch struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Teams       struct {
		Team1 TeamMatchTeam `json:"team1"`
		Team2 TeamMatchTeam `json:"team2"`
	} `json:"teams"`
	Settings  TeamMatchSettings `json:"settings"`
	StartTime int64             `json:"start_time"`
	Boards    int               `json:"boards"`
}

// MatchBoardGame is a game played on a specific match board.
type MatchBoardGame struct {
	Accuracies  *GameAccuracies  `json:"accuracies,omitempty"`
	TimeControl string           `json:"time_control"`
	URL         string           `json:"url"`
	FEN         string           `json:"fen"`
	PGN         string           `json:"pgn"`
	TimeClass   string           `json:"time_class"`
	Rules       string           `json:"rules"`
	ECO         string           `json:"eco,omitempty"`
	Match       string           `json:"match,omitempty"`
	Black       MatchBoardPlayer `json:"black"`
	White       MatchBoardPlayer `json:"white"`
	StartTime   int64            `json:"start_time,omitempty"`
	EndTime     int64            `json:"end_time,omitempty"`
}

// MatchBoardPlayer holds the details for one player on a match board.
type MatchBoardPlayer struct {
	// ID is the API URL of this player's profile.
	ID string `json:"@id"`
	// Username is the player's username.
	Username string `json:"username"`
	// Team is the URL of the player's team club.
	Team string `json:"team,omitempty"`
	// Result is the game result code.
	Result string `json:"result"`
	// Rating is the player's rating.
	Rating int `json:"rating"`
}

// MatchBoard is the response from the team match board endpoint.
type MatchBoard struct {
	// BoardScores maps usernames to their board scores.
	BoardScores map[string]float64 `json:"board_scores"`
	// Games is the list of games played on this board.
	Games []MatchBoardGame `json:"games"`
}

// LiveMatchSettings holds configuration for a live team match.
type LiveMatchSettings struct {
	// Rules is the game variant.
	Rules string `json:"rules"`
	// TimeClass is the time control class.
	TimeClass string `json:"time_class"`
	// TimeControl is the base time in seconds.
	TimeControl int `json:"time_control"`
	// TimeIncrement is the increment in seconds.
	TimeIncrement int `json:"time_increment"`
	// MinTeamPlayers is the minimum number of players per team.
	MinTeamPlayers int `json:"min_team_players"`
	// MinRequiredGames is the minimum number of required games.
	MinRequiredGames int `json:"min_required_games"`
	// Autostart indicates whether the match starts automatically.
	Autostart bool `json:"autostart"`
}

// LiveTeamMatchPlayer is a player entry in a live team match.
type LiveTeamMatchPlayer struct {
	// Username is the player's username.
	Username string `json:"username"`
	// Board is the URL of the player's board.
	Board string `json:"board,omitempty"`
	// Stats is the URL to the player's stats (finished matches).
	Stats string `json:"stats,omitempty"`
	// Status is the account status.
	Status string `json:"status,omitempty"`
	// PlayedAsWhite is the result code for the game played as white.
	PlayedAsWhite string `json:"played_as_white,omitempty"`
	// PlayedAsBlack is the result code for the game played as black.
	PlayedAsBlack string `json:"played_as_black,omitempty"`
}

// LiveTeamMatchTeam holds the details for one side of a live team match.
type LiveTeamMatchTeam struct {
	// ID is the API URL of the club profile.
	ID string `json:"@id"`
	// URL is the web URL of the club profile.
	URL string `json:"url"`
	// Name is the club name.
	Name string `json:"name"`
	// Result is the match result for this team (finished only).
	Result string `json:"result,omitempty"`
	// FairPlayRemovals is the list of usernames removed for fair play violations.
	FairPlayRemovals []string `json:"fair_play_removals"`
	// Players is the list of players on this team.
	Players []LiveTeamMatchPlayer `json:"players"`
	// Score is the team's score.
	Score float64 `json:"score"`
}

// LiveTeamMatch contains details about a live team match.
type LiveTeamMatch struct {
	ID     string `json:"@id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status string `json:"status"`
	Teams  struct {
		Team1 LiveTeamMatchTeam `json:"team1"`
		Team2 LiveTeamMatchTeam `json:"team2"`
	} `json:"teams"`
	Settings  LiveMatchSettings `json:"settings"`
	StartTime int64             `json:"start_time"`
	EndTime   int64             `json:"end_time,omitempty"`
	Boards    int               `json:"boards"`
}

// LiveMatchBoardGame is a game played on a live match board.
type LiveMatchBoardGame struct {
	URL         string               `json:"url"`
	PGN         string               `json:"pgn"`
	TimeControl string               `json:"time_control"`
	FEN         string               `json:"fen"`
	TimeClass   string               `json:"time_class"`
	Rules       string               `json:"rules"`
	ECO         string               `json:"eco,omitempty"`
	White       LiveMatchBoardPlayer `json:"white"`
	Black       LiveMatchBoardPlayer `json:"black"`
	EndTime     int64                `json:"end_time"`
	Rated       bool                 `json:"rated"`
}

// LiveMatchBoardPlayer holds details for one player on a live match board.
type LiveMatchBoardPlayer struct {
	// ID is the API URL of this player's profile.
	ID string `json:"@id"`
	// Username is the player's username.
	Username string `json:"username"`
	// Result is the game result code.
	Result string `json:"result"`
	// Rating is the player's rating.
	Rating int `json:"rating"`
}

// LiveMatchBoard is the response from the live team match board endpoint.
type LiveMatchBoard struct {
	// BoardScores maps usernames to their board scores.
	BoardScores map[string]float64 `json:"board_scores"`
	// Games is the list of games played on this board.
	Games []LiveMatchBoardGame `json:"games"`
}

// Country contains the profile of a country.
type Country struct {
	// ID is the canonical API URL for this country profile.
	ID string `json:"@id"`
	// Name is the human-readable country name.
	Name string `json:"name"`
	// Code is the ISO 3166-1 alpha-2 country code.
	Code string `json:"code"`
}

// CountryPlayers is the response from the country players endpoint.
type CountryPlayers struct {
	// Players is the list of usernames for recently active players in this country.
	Players []string `json:"players"`
}

// CountryClubs is the response from the country clubs endpoint.
type CountryClubs struct {
	// Clubs is the list of API URLs for clubs associated with this country.
	Clubs []string `json:"clubs"`
}

// DailyPuzzle contains information about a daily puzzle.
type DailyPuzzle struct {
	// Title is the puzzle title.
	Title string `json:"title"`
	// URL is the web URL of the puzzle page.
	URL string `json:"url"`
	// FEN is the puzzle's starting FEN position.
	FEN string `json:"fen"`
	// PGN is the puzzle's PGN.
	PGN string `json:"pgn"`
	// Image is the URL of the puzzle image.
	Image string `json:"image"`
	// PublishTime is the Unix timestamp of when the puzzle was published.
	PublishTime int64 `json:"publish_time"`
}

// Streamer contains information about a Chess.com streamer.
type Streamer struct {
	// Username is the streamer's username.
	Username string `json:"username"`
	// Avatar is the URL of the streamer's avatar.
	Avatar string `json:"avatar"`
	// TwitchURL is the streamer's Twitch.tv URL.
	TwitchURL string `json:"twitch_url"`
	// URL is the streamer's Chess.com profile URL.
	URL string `json:"url"`
}

// StreamersList is the response from the streamers endpoint.
type StreamersList struct {
	// Streamers is the list of active Chess.com streamers.
	Streamers []Streamer `json:"streamers"`
}

// LeaderboardEntry is one entry in a leaderboard category.
type LeaderboardEntry struct {
	// ID is the API URL of this player's profile.
	ID string `json:"@id"`
	// URL is the web URL of the player's profile.
	URL string `json:"url"`
	// Username is the player's username.
	Username string `json:"username"`
	// PlayerID is the non-changing Chess.com ID.
	PlayerID int `json:"player_id"`
	// Score is the player's score in this category.
	Score int `json:"score"`
	// Rank is the player's rank (1–50).
	Rank int `json:"rank"`
}

// Leaderboards is the response from the leaderboards endpoint.
type Leaderboards struct {
	Daily             []LeaderboardEntry `json:"daily"`
	Daily960          []LeaderboardEntry `json:"daily960"`
	LiveRapid         []LeaderboardEntry `json:"live_rapid"`
	LiveBlitz         []LeaderboardEntry `json:"live_blitz"`
	LiveBullet        []LeaderboardEntry `json:"live_bullet"`
	LiveBughouse      []LeaderboardEntry `json:"live_bughouse"`
	LiveBlitz960      []LeaderboardEntry `json:"live_blitz960"`
	LiveThreeCheck    []LeaderboardEntry `json:"live_threecheck"`
	LiveCrazyhouse    []LeaderboardEntry `json:"live_crazyhouse"`
	LiveKingOfTheHill []LeaderboardEntry `json:"live_kingofthehill"`
	Lessons           []LeaderboardEntry `json:"lessons"`
	Tactics           []LeaderboardEntry `json:"tactics"`
}
