/**
 * Go Video Downloader
 *
 *    Copyright 2017 Tenta, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For any questions, please contact developer@tenta.io
 *
 * masterregexp.go: Automatically generated regexp by combining URL validation expressions of all extractors
 */

package downloader

func init() {
	masterRegexp = `(?<extr_AnySex>https?://(?:www\.)?anysex\.com/(?:\d+))|(?<extr_ATTTechChannel>https?://techchannel\.att\.com/play-video\.cfm/([^/]+/)*(?:.+))|(?<extr_Bang>https?://(?:www\.)?bang\.com/video/(?:[0-9a-zA-Z]+))|(?<extr_Bild>https?://(?:www\.)?bild\.de/(?:[^/]+/)+(?:[^/]+)-(?:\d+)(?:,auto=true)?\.bild\.html)|(?<extr_Criterion>https?://(?:www\.)?criterion\.com/films/(?:[0-9]+)-.+)|(?<extr_EchoMsk>https?://(?:www\.)?echo\.msk\.ru/sounds/(?:\d+))|(?<extr_FiveTV>http://(?:www\.)?5-tv\.ru/(?:(?:[^/]+/)+(?:\d+)|(?:[^/?#]+)(?:[/?#])?))|(?<extr_Hark>https?://(?:www\.)?hark\.com/clips/(?:.+?)-.+)|(?<extr_HentaiStigma>^https?://hentai\.animestigma\.com/(?:[^/]+))|(?<extr_HistoricFilms>https?://(?:www\.)?historicfilms\.com/(?:tapes/|play)(?:\d+))|(?<extr_Keek>https?://(?:www\.)?keek\.com/keek/(?:\w+))|(?<extr_Morningstar>https?://(?:(?:www|news)\.)morningstar\.com/[cC]over/video[cC]enter\.aspx\?id=(?:[0-9]+))|(?<extr_Slutload>^https?://(?:\w+\.)?slutload\.com/video/[^/]+/(?:[^/]+)/?$)|(?<extr_WDRMobile>https?://mobile-ondemand\.wdr\.de/.*?/fsk(?:[0-9]+)/[0-9]+/[0-9]+/(?:[0-9]+)_(?:[0-9]+))|(?<extr_XHamster>https?://(?:.+?\.)?xhamster\.com/(?:movies/(?:\d+)/(?:[^/]*)\.html|videos/(?:[^/]*)-(?:\d+)))|(?<extr_XNXX>https?://(?:video|www)\.xnxx\.com/video-?(?:[0-9a-z]+)/)|(?<extr_XXXYMovies>https?://(?:www\.)?xxxymovies\.com/videos/(?:\d+)/(?:[^/]+))`
}
