package editor

import (
	"fmt"
)

// SvgXML is
type SvgXML struct {
	xml       string
	width     int
	height    int
	thickness float64
	color     *RGBA
}

// Svg is
type Svg struct {
	width  int
	height int
	color  *RGBA
	bg     *RGBA
	name   string
}

func (e *Editor) getSvgs() map[string]*SvgXML {
	e.svgsOnce.Do(func() {
		e.initSVGS()
	})
	return e.svgs
}

func (e *Editor) getSvg(name string, color *RGBA) string {
	e.svgsOnce.Do(func() {
		e.initSVGS()
	})
	svg := e.svgs[name]
	if svg == nil {
		svg = e.svgs["default"]
	}
	if color == nil {
		color = svg.color
	}
	if color == nil {
		color = e.fgcolor
	}
	if color == nil {
		color = newRGBA(255, 255, 255, 1)
	}

	return fmt.Sprintf(svg.xml, color.Hex())
}

func (e *Editor) initSVGS() {
	blue := newRGBA(81, 154, 186, 1)
	e.svgs = map[string]*SvgXML{}

	e.svgs["chevron-down"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M7.41,8.58L12,13.17L16.59,8.58L18,10L12,16L6,10L7.41,8.58Z" /></svg>`,
	}

	e.svgs["chevron-right"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z" /></svg>`,
	}

	e.svgs["bell"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M21,19V20H3V19L5,17V11C5,7.9 7.03,5.17 10,4.29C10,4.19 10,4.1 10,4A2,2 0 0,1 12,2A2,2 0 0,1 14,4C14,4.1 14,4.19 14,4.29C16.97,5.17 19,7.9 19,11V17L21,19M14,21A2,2 0 0,1 12,23A2,2 0 0,1 10,21" /></svg>`,
	}

	e.svgs["warn"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M13,14H11V10H13M13,18H11V16H13M1,21H23L12,2L1,21Z" /></svg>`,
	}

	e.svgs["info"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M13,9H11V7H13M13,17H11V11H13M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z" /></svg>`,
	}

	e.svgs["progress"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><g transform="translate(0,2)"><path fill="%s" d="M13 2.03V2.05L13 4.05C17.39 4.59 20.5 8.58 19.96 12.97C19.5 16.61 16.64 19.5 13 19.93V21.93C18.5 21.38 22.5 16.5 21.95 11C21.5 6.25 17.73 2.5 13 2.03M11 2.06C9.05 2.25 7.19 3 5.67 4.26L7.1 5.74C8.22 4.84 9.57 4.26 11 4.06V2.06M4.26 5.67C3 7.19 2.25 9.04 2.05 11H4.05C4.24 9.58 4.8 8.23 5.69 7.1L4.26 5.67M2.06 13C2.26 14.96 3.03 16.81 4.27 18.33L5.69 16.9C4.81 15.77 4.24 14.42 4.06 13H2.06M7.1 18.37L5.67 19.74C7.18 21 9.04 21.79 11 22V20C9.58 19.82 8.23 19.25 7.1 18.37M12.5 7V12.25L17 14.92L16.25 16.15L11 13V7H12.5Z" /></g></svg>`,
	}

	e.svgs["settings"] = &SvgXML{
		width:  128,
		height: 128,
		xml:    `<svg width="24" height="24" viewBox="0 0 16 16"><path fill="%s" d="M14 8.77v-1.6l-1.94-.64-.45-1.09.88-1.84-1.13-1.13-1.81.91-1.09-.45-.69-1.92h-1.6l-.63 1.94-1.11.45-1.84-.88-1.13 1.13.91 1.81-.45 1.09L0 7.23v1.59l1.94.64.45 1.09-.88 1.84 1.13 1.13 1.81-.91 1.09.45.69 1.92h1.59l.63-1.94 1.11-.45 1.84.88 1.13-1.13-.92-1.81.47-1.09L14 8.75v.02zM7 11c-1.66 0-3-1.34-3-3s1.34-3 3-3 3 1.34 3 3-1.34 3-3 3z"></path></svg>`,
	}

	e.svgs["puzzle"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M20.5 11H19V7C19 5.89 18.1 5 17 5H13V3.5A2.5 2.5 0 0 0 10.5 1A2.5 2.5 0 0 0 8 3.5V5H4A2 2 0 0 0 2 7V10.8H3.5C5 10.8 6.2 12 6.2 13.5C6.2 15 5 16.2 3.5 16.2H2V20A2 2 0 0 0 4 22H7.8V20.5C7.8 19 9 17.8 10.5 17.8C12 17.8 13.2 19 13.2 20.5V22H17A2 2 0 0 0 19 20V16H20.5A2.5 2.5 0 0 0 23 13.5A2.5 2.5 0 0 0 20.5 11Z" /></svg>`,
	}

	e.svgs["timer"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M12 20A7 7 0 0 1 5 13A7 7 0 0 1 12 6A7 7 0 0 1 19 13A7 7 0 0 1 12 20M19.03 7.39L20.45 5.97C20 5.46 19.55 5 19.04 4.56L17.62 6C16.07 4.74 14.12 4 12 4A9 9 0 0 0 3 13A9 9 0 0 0 12 22C17 22 21 17.97 21 13C21 10.88 20.26 8.93 19.03 7.39M11 14H13V8H11M15 1H9V3H15V1Z" /></svg>`,
	}

	e.svgs["configfile"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="256" height="256" viewBox="0 0 16 16" version="1.1" aria-hidden="true"><path fill="%s" d="M11.41 9H.59C0 9 0 8.59 0 8c0-.59 0-1 .59-1H11.4c.59 0 .59.41.59 1 0 .59 0 1-.59 1h.01zm0-4H.59C0 5 0 4.59 0 4c0-.59 0-1 .59-1H11.4c.59 0 .59.41.59 1 0 .59 0 1-.59 1h.01zM.59 11H11.4c.59 0 .59.41.59 1 0 .59 0 1-.59 1H.59C0 13 0 12.59 0 12c0-.59 0-1 .59-1z"></path></svg>`,
	}

	e.svgs["download"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="256" height="256" viewBox="0 0 16 16"><g transform="scale(0.88) translate(0,2)"><path fill="%s" d="M9 12h2l-3 3-3-3h2V7h2v5zm3-8c0-.44-.91-3-4.5-3C5.08 1 3 2.92 3 5 1.02 5 0 6.52 0 8c0 1.53 1 3 3 3h3V9.7H3C1.38 9.7 1.3 8.28 1.3 8c0-.17.05-1.7 1.7-1.7h1.3V5c0-1.39 1.56-2.7 3.2-2.7 2.55 0 3.13 1.55 3.2 1.8v1.2H12c.81 0 2.7.22 2.7 2.2 0 2.09-2.25 2.2-2.7 2.2h-2V11h2c2.08 0 4-1.16 4-3.5C16 5.06 14.08 4 12 4z"></path></g></svg>`,
	}

	e.svgs["downloaded"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><g transform="translate(0,2)"><path fill="%s" d="M5 20H19V18H5M19 9H15V3H9V9H5L12 16L19 9Z" /></g></svg>`,
	}

	e.svgs["user"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><g transform="translate(0,2)"><path fill="%s" d="M12 4A4 4 0 0 1 16 8A4 4 0 0 1 12 12A4 4 0 0 1 8 8A4 4 0 0 1 12 4M12 14C16.42 14 20 15.79 20 18V20H4V18C4 15.79 7.58 14 12 14Z" /></g></svg>`,
	}

	e.svgs["star"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><g transform="translate(0,1)"><path fill="%s" d="M12 17.27L18.18 21L16.54 13.97L22 9.24L14.81 8.62L12 2L9.19 8.62L2 9.24L7.45 13.97L5.82 21L12 17.27Z" /></g></svg>`,
	}

	e.svgs["activityedit"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M6 2A2 2 0 0 0 4 4V20A2 2 0 0 0 6 22H18A2 2 0 0 0 20 20V8L14 2H6M6 4H13V9H18V20H6V4M8 12V14H16V12H8M8 16V18H13V16H8Z" /></svg>`,
	}

	e.svgs["activitydein"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M7 2V13H10V22L17 10H13L17 2H7Z" /></svg>`,
	}

	e.svgs["moredots"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M16 12A2 2 0 0 1 18 10A2 2 0 0 1 20 12A2 2 0 0 1 18 14A2 2 0 0 1 16 12M10 12A2 2 0 0 1 12 10A2 2 0 0 1 14 12A2 2 0 0 1 12 14A2 2 0 0 1 10 12M4 12A2 2 0 0 1 6 10A2 2 0 0 1 8 12A2 2 0 0 1 6 14A2 2 0 0 1 4 12Z" /></svg>`,
	}

	e.svgs["circle"] = &SvgXML{
		width:  24,
		height: 24,
		xml:    `<svg width="24px" height="24px" viewBox="0 0 24 24"><g transform="scale(0.9) translate(0,3)"><path fill="%s" d="M12 2A10 10 0 0 0 2 12A10 10 0 0 0 12 22A10 10 0 0 0 22 12A10 10 0 0 0 12 2Z" /></g></svg>`,
	}

	e.svgs["directory"] = &SvgXML{
		width:  24,
		height: 24,
		xml: `<?xml version="1.0" encoding="utf-8"?>
<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="%s" d="M10 4H4C2.89 4 2 4.89 2 6V18A2 2 0 0 0 4 20H20A2 2 0 0 0 22 18V8C22 6.89 21.1 6 20 6H12L10 4Z" /></svg>`,
	}

	e.svgs["command"] = &SvgXML{
		width:  24,
		height: 24,
		color:  e.fgcolor,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><g transform="scale(0.93)"><path fill="%s" d="M6 2A4 4 0 0 1 10 6V8H14V6A4 4 0 0 1 18 2A4 4 0 0 1 22 6A4 4 0 0 1 18 10H16V14H18A4 4 0 0 1 22 18A4 4 0 0 1 18 22A4 4 0 0 1 14 18V16H10V18A4 4 0 0 1 6 22A4 4 0 0 1 2 18A4 4 0 0 1 6 14H8V10H6A4 4 0 0 1 2 6A4 4 0 0 1 6 2M16 18A2 2 0 0 0 18 20A2 2 0 0 0 20 18A2 2 0 0 0 18 16H16V18M14 10H10V14H14V10M6 16A2 2 0 0 0 4 18A2 2 0 0 0 6 20A2 2 0 0 0 8 18V16H6M8 6A2 2 0 0 0 6 4A2 2 0 0 0 4 6A2 2 0 0 0 6 8H8V6M18 8A2 2 0 0 0 20 6A2 2 0 0 0 18 4A2 2 0 0 0 16 6V8H18Z" /></g></svg>`,
	}

	e.svgs["replace"] = &SvgXML{
		width:  24,
		height: 24,
		color:  e.fgcolor,
		xml:    `<svg style="width:24px;height:24px" viewBox="0 0 24 24"><path fill="%s" d="M12 5C16.97 5 21 7.69 21 11C21 12.68 19.96 14.2 18.29 15.29C19.36 14.42 20 13.32 20 12.13C20 9.29 16.42 7 12 7V10L8 6L12 2V5M12 19C7.03 19 3 16.31 3 13C3 11.32 4.04 9.8 5.71 8.71C4.64 9.58 4 10.68 4 11.88C4 14.71 7.58 17 12 17V14L16 18L12 22V19Z" /></svg>`,
	}

	// terminal
	e.svgs["terminal"] = &SvgXML{
		width:  112,
		height: 128,
		xml:    `<svg height="128" viewBox="0 0 14 16" version="1.1" width="112" aria-hidden="true"><path fill="%s" fill-rule="evenodd" d="M7 10h4v1H7v-1zm-3 1l3-3-3-3-.75.75L5.5 8l-2.25 2.25L4 11zm10-8v10c0 .55-.45 1-1 1H1c-.55 0-1-.45-1-1V3c0-.55.45-1 1-1h12c.55 0 1 .45 1 1zm-1 0H1v10h12V3z"></path></svg>`,
	}

	// for Insert Mode
	e.svgs["edit"] = &SvgXML{
		width:  24,
		height: 24,
		color:  e.fgcolor,
		xml: `<?xml version="1.0" encoding="utf-8"?>
   <svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><g transform="translate(0,2)"><path fill="%s" d="M20.71 4.04C21.1 3.65 21.1 3 20.71 2.63L18.37 0.29C18 -0.1 17.35 -0.1 16.96 0.29L15 2.25L18.75 6M17.75 7L14 3.25L4 13.25V17H7.75L17.75 7Z" /></g></svg>`,
	}

	// for Visual Mode
	e.svgs["select"] = &SvgXML{
		width:  24,
		height: 24,
		color:  e.fgcolor,
		xml: `<?xml version="1.0" encoding="utf-8"?>
<svg width="24" height="24" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path fill="%s" d="M4 3H5V5H3V4A1 1 0 0 1 4 3M20 3A1 1 0 0 1 21 4V5H19V3H20M15 5V3H17V5H15M11 5V3H13V5H11M7 5V3H9V5H7M21 20A1 1 0 0 1 20 21H19V19H21V20M15 21V19H17V21H15M11 21V19H13V21H11M7 21V19H9V21H7M4 21A1 1 0 0 1 3 20V19H5V21H4M3 15H5V17H3V15M21 15V17H19V15H21M3 11H5V13H3V11M21 11V13H19V11H21M3 7H5V9H3V7M21 7V9H19V7H21Z" /></svg>`,
	}

	e.svgs["fire"] = &SvgXML{
		width:  1792,
		height: 1792,
		xml: `<?xml version="1.0" encoding="utf-8"?>
<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><path fill="%s" d="M1600 1696v64q0 13-9.5 22.5t-22.5 9.5h-1344q-13 0-22.5-9.5t-9.5-22.5v-64q0-13 9.5-22.5t22.5-9.5h1344q13 0 22.5 9.5t9.5 22.5zm-256-1056q0 78-24.5 144t-64 112.5-87.5 88-96 77.5-87.5 72-64 81.5-24.5 96.5q0 96 67 224l-4-1 1 1q-90-41-160-83t-138.5-100-113.5-122.5-72.5-150.5-27.5-184q0-78 24.5-144t64-112.5 87.5-88 96-77.5 87.5-72 64-81.5 24.5-96.5q0-94-66-224l3 1-1-1q90 41 160 83t138.5 100 113.5 122.5 72.5 150.5 27.5 184z"/></svg>`,
	}
	e.svgs["comment"] = &SvgXML{
		width:  1792,
		height: 1792,
		color:  blue,
		xml: `<?xml version="1.0" encoding="utf-8"?>
<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><path fill="%s" d="M1792 896q0 174-120 321.5t-326 233-450 85.5q-70 0-145-8-198 175-460 242-49 14-114 22-17 2-30.5-9t-17.5-29v-1q-3-4-.5-12t2-10 4.5-9.5l6-9 7-8.5 8-9q7-8 31-34.5t34.5-38 31-39.5 32.5-51 27-59 26-76q-157-89-247.5-220t-90.5-281q0-130 71-248.5t191-204.5 286-136.5 348-50.5q244 0 450 85.5t326 233 120 321.5z"/></svg>`,
	}
	e.svgs["linterr"] = &SvgXML{
		width:     24,
		height:    24,
		thickness: 2,
		color:     e.fgcolor,
		xml:       `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M13 13H11V7H13M13 17H11V15H13M12 2A10 10 0 0 0 2 12A10 10 0 0 0 12 22A10 10 0 0 0 22 12A10 10 0 0 0 12 2Z" /></svg>`,
	}
	e.svgs["lintwrn"] = &SvgXML{
		width:     24,
		height:    24,
		thickness: 2,
		color:     e.fgcolor,
		xml:       `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M13 14H11V10H13M13 18H11V16H13M1 21H23L12 2L1 21Z" /></svg>`,
	}
	e.svgs["hoverclose"] = &SvgXML{
		width:  1792,
		height: 1792,
		color:  e.fgcolor,
		xml:    `<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><path fill="%s" d="M1277 1122q0-26-19-45l-181-181 181-181q19-19 19-45 0-27-19-46l-90-90q-19-19-46-19-26 0-45 19l-181 181-181-181q-19-19-45-19-27 0-46 19l-90 90q-19 19-19 46 0 26 19 45l181 181-181 181q-19 19-19 45 0 27 19 46l90 90q19 19 46 19 26 0 45-19l181-181 181 181q19 19 45 19 27 0 46-19l90-90q19-19 19-46zm387-226q0 209-103 385.5t-279.5 279.5-385.5 103-385.5-103-279.5-279.5-103-385.5 103-385.5 279.5-279.5 385.5-103 385.5 103 279.5 279.5 103 385.5z"/></svg>`,
	}
	e.svgs["cross"] = &SvgXML{
		width:     100,
		height:    100,
		thickness: 2,
		color:     e.fgcolor,
		xml:       `<?xml version="1.0" encoding="utf-8"?><svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><path fill="%s" d="M1490 1322q0 40-28 68l-136 136q-28 28-68 28t-68-28l-294-294-294 294q-28 28-68 28t-68-28l-136-136q-28-28-28-68t28-68l294-294-294-294q-28-28-28-68t28-68l136-136q28-28 68-28t68 28l294 294 294-294q28-28 68-28t68 28l136 136q28 28 28 68t-28 68l-294 294 294 294q28 28 28 68z"/></svg>`,
	}
	e.svgs["bad"] = &SvgXML{
		width:  24,
		height: 24,
		color:  e.fgcolor,
		xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M11 15H13V17H11V15M11 7H13V13H11V7M12 2C6.47 2 2 6.5 2 12A10 10 0 0 0 12 22A10 10 0 0 0 22 12A10 10 0 0 0 12 2M12 20A8 8 0 0 1 4 12A8 8 0 0 1 12 4A8 8 0 0 1 20 12A8 8 0 0 1 12 20Z" /></svg>`,
		//xml:    `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M13.46 12L19 17.54V19H17.54L12 13.46L6.46 19H5V17.54L10.54 12L5 6.46V5H6.46L12 10.54L17.54 5H19V6.46L13.46 12Z" /></svg>`,
	}

	e.svgs["default"] = &SvgXML{
		width:     200,
		height:    200,
		thickness: 0.5,
		color:     e.fgcolor,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 200"><g><path fill="%s" d="M22.34999879828441,17.7129974066314 h147.91225647948028 v20.748584330809138 H22.34999879828441 V17.7129974066314 zM22.34999879828441,65.18324337560382 h126.22055467908892 v20.748584330809138 H22.34999879828441 V65.18324337560382 zM22.34999879828441,113.91097930401922 h155.3000099912078 v20.748584330809138 H22.34999879828441 V113.91097930401922 zM22.34999879828441,161.538411517922 h91.01083581468554 v20.748584330809138 H22.34999879828441 V161.538411517922 z"/></g></svg>`,
	}
	e.svgs["empty"] = &SvgXML{
		width:  200,
		height: 200,
		color:  e.fgcolor,
		xml:    `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 200 200"><g><path fill="%s" d=""/></g></svg>`,
	}
	e.svgs["folder"] = &SvgXML{
		width:     1792,
		height:    1792,
		thickness: 0.5,
		color:     e.fgcolor,
		xml:       `<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><g><path fill="%s" d="M1600 1312v-704q0-40-28-68t-68-28h-704q-40 0-68-28t-28-68v-64q0-40-28-68t-68-28h-320q-40 0-68 28t-28 68v960q0 40 28 68t68 28h1216q40 0 68-28t28-68zm128-704v704q0 92-66 158t-158 66h-1216q-92 0-158-66t-66-158v-960q0-92 66-158t158-66h320q92 0 158 66t66 158v32h672q92 0 158 66t66 158z"/></g></svg>`,
	}
	e.svgs["git"] = &SvgXML{
		width:  1024,
		height: 1024,
		color:  e.fgcolor,
		xml:    `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-90 -40 1050 1024" height="1024" width="640"><g transform="scale(0.95)"><path fill="%s" d="M512 192c-71 0-128 57-128 128 0 47 26 88 64 110v18c0 64-64 128-128 128-53 0-95 11-128 29v-303c38-22 64-63 64-110 0-71-57-128-128-128s-128 57-128 128c0 47 26 88 64 110v419c-38 22-64 63-64 110 0 71 57 128 128 128s128-57 128-128c0-34-13-64-34-87 19-23 49-41 98-41 128 0 256-128 256-256v-18c38-22 64-63 64-110 0-71-57-128-128-128z m-384-64c35 0 64 29 64 64s-29 64-64 64-64-29-64-64 29-64 64-64z m0 768c-35 0-64-29-64-64s29-64 64-64 64 29 64 64-29 64-64 64z m384-512c-35 0-64-29-64-64s29-64 64-64 64 29 64 64-29 64-64 64z"/></g></svg>`,
		//		xml:       `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M2.6 10.59L8.38 4.8L10.07 6.5C9.83 7.35 10.22 8.28 11 8.73V14.27C10.4 14.61 10 15.26 10 16A2 2 0 0 0 12 18A2 2 0 0 0 14 16C14 15.26 13.6 14.61 13 14.27V9.41L15.07 11.5C15 11.65 15 11.82 15 12A2 2 0 0 0 17 14A2 2 0 0 0 19 12A2 2 0 0 0 17 10C16.82 10 16.65 10 16.5 10.07L13.93 7.5C14.19 6.57 13.71 5.55 12.78 5.16C12.35 5 11.9 4.96 11.5 5.07L9.8 3.38L10.59 2.6C11.37 1.81 12.63 1.81 13.41 2.6L21.4 10.59C22.19 11.37 22.19 12.63 21.4 13.41L13.41 21.4C12.63 22.19 11.37 22.19 10.59 21.4L2.6 13.41C1.81 12.63 1.81 11.37 2.6 10.59Z" /></svg>`,
	}
	e.svgs["check"] = &SvgXML{
		width:     1792,
		height:    1792,
		thickness: 0,
		color:     newRGBA(141, 193, 73, 1),
		xml:       `<svg width="1792" height="1792" viewBox="0 0 1792 1792" xmlns="http://www.w3.org/2000/svg"><g><path fill="%s" d="M1671 566q0 40-28 68l-724 724-136 136q-28 28-68 28t-68-28l-136-136-362-362q-28-28-28-68t28-68l136-136q28-28 68-28t68 28l294 295 656-657q28-28 68-28t68 28l136 136q28 28 28 68z"/></g></svg>`,
	}
	e.svgs["exclamation"] = &SvgXML{
		width:     24,
		height:    24,
		thickness: 0,
		color:     e.fgcolor,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M12 2L1 21H23M12 6L19.53 19H4.47M11 10V14H13V10M11 16V18H13V16" /></svg>`,
	}

	e.svgs["json"] = &SvgXML{
		width:     24,
		height:    24,
		thickness: 0,
		color:     blue,
		xml:       `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M5 3H7V5H5V10A2 2 0 0 1 3 12A2 2 0 0 1 5 14V19H7V21H5C3.93 20.73 3 20.1 3 19V15A2 2 0 0 0 1 13H0V11H1A2 2 0 0 0 3 9V5A2 2 0 0 1 5 3M19 3A2 2 0 0 1 21 5V9A2 2 0 0 0 23 11H24V13H23A2 2 0 0 0 21 15V19A2 2 0 0 1 19 21H17V19H19V14A2 2 0 0 1 21 12A2 2 0 0 1 19 10V5H17V3H19M12 15A1 1 0 0 1 13 16A1 1 0 0 1 12 17A1 1 0 0 1 11 16A1 1 0 0 1 12 15M8 15A1 1 0 0 1 9 16A1 1 0 0 1 8 17A1 1 0 0 1 7 16A1 1 0 0 1 8 15M16 15A1 1 0 0 1 17 16A1 1 0 0 1 16 17A1 1 0 0 1 15 16A1 1 0 0 1 16 15Z" /></svg>`,
	}

	e.svgs["ts"] = &SvgXML{
		width:     24,
		height:    24,
		thickness: 0,
		color:     blue,
		xml:       `<svg width="24" height="24" viewBox="0 0 24 24"><path fill="%s" d="M3 3H21V21H3V3M13.71 17.86C14.21 18.84 15.22 19.59 16.8 19.59C18.4 19.59 19.6 18.76 19.6 17.23C19.6 15.82 18.79 15.19 17.35 14.57L16.93 14.39C16.2 14.08 15.89 13.87 15.89 13.37C15.89 12.96 16.2 12.64 16.7 12.64C17.18 12.64 17.5 12.85 17.79 13.37L19.1 12.5C18.55 11.54 17.77 11.17 16.7 11.17C15.19 11.17 14.22 12.13 14.22 13.4C14.22 14.78 15.03 15.43 16.25 15.95L16.67 16.13C17.45 16.47 17.91 16.68 17.91 17.26C17.91 17.74 17.46 18.09 16.76 18.09C15.93 18.09 15.45 17.66 15.09 17.06L13.71 17.86M13 11.25H8V12.75H9.5V20H11.25V12.75H13V11.25Z" /></svg>`,
	}

	e.svgs["js"] = &SvgXML{
		width:     24,
		height:    24,
		thickness: 0,
		color:     blue,
		xml:       `<svg style="width:24px;height:24px" viewBox="0 0 24 24"><path fill="%s" d="M3 3H21V21H3V3M7.73 18.04C8.13 18.89 8.92 19.59 10.27 19.59C11.77 19.59 12.8 18.79 12.8 17.04V11.26H11.1V17C11.1 17.86 10.75 18.08 10.2 18.08C9.62 18.08 9.38 17.68 9.11 17.21L7.73 18.04M13.71 17.86C14.21 18.84 15.22 19.59 16.8 19.59C18.4 19.59 19.6 18.76 19.6 17.23C19.6 15.82 18.79 15.19 17.35 14.57L16.93 14.39C16.2 14.08 15.89 13.87 15.89 13.37C15.89 12.96 16.2 12.64 16.7 12.64C17.18 12.64 17.5 12.85 17.79 13.37L19.1 12.5C18.55 11.54 17.77 11.17 16.7 11.17C15.19 11.17 14.22 12.13 14.22 13.4C14.22 14.78 15.03 15.43 16.25 15.95L16.67 16.13C17.45 16.47 17.91 16.68 17.91 17.26C17.91 17.74 17.46 18.09 16.76 18.09C15.93 18.09 15.45 17.66 15.09 17.06L13.71 17.86Z" /></svg>`,
	}

	e.svgs["markdown"] = &SvgXML{
		width:     1024,
		height:    1024,
		thickness: 0,
		color:     blue,
		xml:       `<svg height="1024" width="1024" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"><path fill="%s" d="M950.154 192H73.846C33.127 192 0 225.12699999999995 0 265.846v492.308C0 798.875 33.127 832 73.846 832h876.308c40.721 0 73.846-33.125 73.846-73.846V265.846C1024 225.12699999999995 990.875 192 950.154 192zM576 703.875L448 704V512l-96 123.077L256 512v192H128V320h128l96 128 96-128 128-0.125V703.875zM767.091 735.875L608 512h96V320h128v192h96L767.091 735.875z" /></svg>`,
	}

	e.svgs["sh"] = &SvgXML{
		width:     18,
		height:    18,
		thickness: 0,
		color:     e.fgcolor,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 18 18"><g><path fill="%s" d="M1.9426860809326185,2.020794440799331 V15.979206420212122 H16.057314060551434 V2.020794440799331 zm2.1962214082875993,1.8041484023191938 l1.488652194536292,1.1762194099782244 l-1.488652194536292,1.2543270180082322 zm0,3.606765445891675 H13.234695018669578 v1.0200022335531784 H4.138907489220218 zm0,2.9007278261308826 h5.723346475939733 v1.0200022335531784 H4.138907489220218 zm0,2.824150479043169 h7.60560350567875 v1.018470796391866 H4.138907489220218 z" /></g></svg>`,
	}
	e.svgs["py"] = &SvgXML{
		width:     128,
		height:    128,
		thickness: 0,
		color:     blue,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128"><g><path fill="%s" d="M49.33 62h29.159c8.117 0 14.511-6.868 14.511-15.019v-27.798c0-7.912-6.632-13.856-14.555-15.176-5.014-.835-10.195-1.215-15.187-1.191-4.99.023-9.612.448-13.805 1.191-12.355 2.181-14.453 6.751-14.453 15.176v10.817h29v4h-40.224000000000004c-8.484 0-15.914 5.108-18.237 14.811-2.681 11.12-2.8 17.919 0 29.53 2.075 8.642 7.03 14.659 15.515 14.659h9.946v-13.048c0-9.637 8.428-17.952 18.33-17.952zm-1.838-39.11c-3.026 0-5.478-2.479-5.478-5.545 0-3.079 2.451-5.581 5.478-5.581 3.015 0 5.479 2.502 5.479 5.581-.001 3.066-2.465 5.545-5.479 5.545zM122.281 48.811c-2.098-8.448-6.103-14.811-14.599-14.811h-10.682v12.981c0 10.05-8.794 18.019-18.511 18.019h-29.159c-7.988 0-14.33 7.326-14.33 15.326v27.8c0 7.91 6.745 12.564 14.462 14.834 9.242 2.717 17.994 3.208 29.051 0 7.349-2.129 14.487-6.411 14.487-14.834v-11.126h-29v-4h43.682c8.484 0 11.647-5.776 14.599-14.66 3.047-9.145 2.916-17.799 0-29.529zm-41.955 55.606c3.027 0 5.479 2.479 5.479 5.547 0 3.076-2.451 5.579-5.479 5.579-3.015 0-5.478-2.502-5.478-5.579 0-3.068 2.463-5.547 5.478-5.547z"/></g></svg>`,
	}
	e.svgs["pyc"] = e.svgs["py"]
	e.svgs["c"] = &SvgXML{
		width:     128,
		height:    128,
		thickness: 1,
		color:     blue,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128"><g><path fill="%s" d="M97.5221405461895,45.03725476832212 L108.76640973908529,40.28547464292772 C108.56169322951193,38.08885210339795 94.3709708336648,16.355134752322773 69.08566426354695,16.60775090230729 C43.80035769342909,16.86037464885613 18.857250847109363,37.373658362115926 19.237894612168745,64.2531288574854 C19.61853837722812,91.1325841597262 44.09328059332486,111.39442159390717 68.83481490446007,111.39442919047148 C93.57635717502413,111.39443678703581 108.34661354359856,92.47813321193519 107.91159096075516,92.0739200243284 C107.47656837791175,91.66970683672162 100.93955315316323,85.45629455970422 100.26450603444523,85.0460952794112 C99.58945891572723,84.63589599911819 90.07591976589798,98.4070389619417 68.86888921929224,98.15440761882854 C47.66185867268652,97.90177627571538 33.5390540830676,80.57280520208417 33.375925589144344,64.50575260403423 C33.212797095221084,48.43870000598427 47.89705183525975,31.110936786080387 69.10406646300783,30.858313039531545 C90.31109700961355,30.605681696418387 93.33611773457105,42.36367434123211 97.5221405461895,45.03725476832212 z"></path></g></svg>`,
	}
	e.svgs["cpp"] = &SvgXML{
		width:     24,
		height:    24,
		thickness: 0.5,
		color:     blue,
		xml: `<svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
    <path fill="%s" d="M10.5 15.97L10.91 18.41C10.65 18.55 10.23 18.68 9.67 18.8C9.1 18.93 8.43 19 7.66 19C5.45 18.96 3.79 18.3 2.68 17.04C1.56 15.77 1 14.16 1 12.21C1.05 9.9 1.72 8.13 3 6.89C4.32 5.64 5.96 5 7.94 5C8.69 5 9.34 5.07 9.88 5.19C10.42 5.31 10.82 5.44 11.08 5.59L10.5 8.08L9.44 7.74C9.04 7.64 8.58 7.59 8.05 7.59C6.89 7.58 5.93 7.95 5.18 8.69C4.42 9.42 4.03 10.54 4 12.03C4 13.39 4.37 14.45 5.08 15.23C5.79 16 6.79 16.4 8.07 16.41L9.4 16.29C9.83 16.21 10.19 16.1 10.5 15.97M11 11H13V9H15V11H17V13H15V15H13V13H11V11M18 11H20V9H22V11H24V13H22V15H20V13H18V11Z" />
</svg>`,
	}
	e.svgs["go"] = &SvgXML{
		width:     128,
		height:    128,
		thickness: 0,
		color:     blue,
		xml:       `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 128 128"><g><path fill="%s" d="M108.2 64.8c-.1-.1-.2-.2-.4-.2l-.1-.1c-.1-.1-.2-.1-.2-.2l-.1-.1c-.1 0-.2-.1-.2-.1l-.2-.1c-.1 0-.2-.1-.2-.1l-.2-.1c-.1 0-.2-.1-.2-.1-.1 0-.1 0-.2-.1l-.3-.1c-.1 0-.1 0-.2-.1l-.3-.1h-.1l-.4-.1h-.2c-.1 0-.2 0-.3-.1h-2.3c-.6-13.3.6-26.8-2.8-39.6 12.9-4.6 2.8-22.3-8.4-14.4-7.4-6.4-17.6-7.8-28.3-7.8-10.5.7-20.4 2.9-27.4 8.4-2.8-1.4-5.5-1.8-7.9-1.1v.1c-.1 0-.3.1-.4.2-.1 0-.3.1-.4.2h-.1c-.1 0-.2.1-.4.2h-.1l-.3.2h-.1l-.3.2h-.1l-.3.2s-.1 0-.1.1l-.3.2s-.1 0-.1.1l-.3.2s-.1 0-.1.1l-.3.2-.1.1c-.1.1-.2.1-.2.2l-.1.1-.2.2-.1.1c-.1.1-.1.2-.2.2l-.1.1c-.1.1-.1.2-.2.2l-.1.1c-.1.1-.1.2-.2.2l-.1.1c-.1.1-.1.2-.2.2l-.1.1c-.1.1-.1.2-.2.2l-.1.1-.1.3s0 .1-.1.1l-.1.3s0 .1-.1.1l-.1.3s0 .1-.1.1l-.1.3s0 .1-.1.1c.4.3.4.4.4.4v.1l-.1.3v.1c0 .1 0 .2-.1.3v3.1c0 .1 0 .2.1.3v.1l.1.3v.1l.1.3s0 .1.1.1l.1.3s0 .1.1.1l.1.3s0 .1.1.1l.2.3s0 .1.1.1l.2.3s0 .1.1.1l.2.3.1.1.3.3.3.3h.1c1 .9 2 1.6 4 2.2v-.2c-4.2 12.6-.7 25.3-.5 38.3-.6 0-.7.4-1.7.5h-.5c-.1 0-.3 0-.5.1-.1 0-.3 0-.4.1l-.4.1h-.1l-.4.1h-.1l-.3.1h-.1l-.3.1s-.1 0-.1.1l-.3.1-.2.1c-.1 0-.2.1-.2.1l-.2.1-.2.1c-.1 0-.2.1-.2.1l-.2.1-.4.3c-.1.1-.2.2-.3.2l-.4.4-.1.1c-.1.2-.3.4-.4.5l-.2.3-.3.6-.1.3v.3c0 .5.2.9.9 1.2.2 3.7 3.9 2 5.6.8l.1-.1c.2-.2.5-.3.6-.3h.1l.2-.1c.1 0 .1 0 .2-.1.2-.1.4-.1.5-.2.1 0 .1-.1.1-.2l.1-.1c.1-.2.2-.6.2-1.2l.1-1.3v1.8c-.5 13.1-4 30.7 3.3 42.5 1.3 2.1 2.9 3.9 4.7 5.4h-.5c-.2.2-.5.4-.8.6l-.9.6-.3.2-.6.4-.9.7-1.1 1c-.2.2-.3.4-.4.5l-.4.6-.2.3c-.1.2-.2.4-.2.6l-.1.3c-.2.8 0 1.7.6 2.7l.4.4h.2c.1 0 .2 0 .4.1.2.4 1.2 2.5 3.9.9 2.8-1.5 4.7-4.6 8.1-5.1l-.5-.6c5.9 2.8 12.8 4 19 4.2 8.7.3 18.6-.9 26.5-5.2 2.2.7 3.9 3.9 5.8 5.4l.1.1.1.1.1.1.1.1s.1 0 .1.1c0 0 .1 0 .1.1 0 0 .1 0 .1.1h2.1000000000000005s.1 0 .1-.1h.1s.1 0 .1-.1h.1s.1 0 .1-.1c0 0 .1 0 .1-.1l.1-.1s.1 0 .1-.1l.1-.1h.1l.2-.2.2-.1h.1l.1-.1h.1l.1-.1.1-.1.1-.1.1-.1.1-.1.1-.1.1-.1v-.1s0-.1.1-.1v-.1s0-.1.1-.1v-.1s0-.1.1-.1v-1.4000000000000001s-.3 0-.3-.1l-.3-.1v-.1l.3-.1s.2 0 .2-.1l.1-.1v-2.1000000000000005s0-.1-.1-.1v-.1s0-.1-.1-.1v-.1s0-.1-.1-.1c0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1 0 0 0-.1-.1-.1l-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1v-.1l-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1-.1c2-1.9 3.8-4.2 5.1-6.9 5.9-11.8 4.9-26.2 4.1-39.2h.1c.1 0 .2.1.2.1h.30000000000000004s.1 0 .1.1h.1s.1 0 .1.1l.2.1c1.7 1.2 5.4 2.9 5.6-.8 1.6.6-.3-1.8-1.3-2.5zm-72.2-41.8c-3.2-16 22.4-19 23.3-3.4.8 13-20 16.3-23.3 3.4zm36.1 15c-1.3 1.4-2.7 1.2-4.1.7 0 1.9.4 3.9.1 5.9-.5.9-1.5 1-2.3 1.4-1.2-.2-2.1-.9-2.6-2l-.2-.1c-3.9 5.2-6.3-1.1-5.2-5-1.2.1-2.2-.2-3-1.5-1.4-2.6.7-5.8 3.4-6.3.7 3 8.7 2.6 10.1-.2 3.1 1.5 6.5 4.3 3.8 7.1zm-7-17.5c-.9-13.8 20.3-17.5 23.4-4 3.5 15-20.8 18.9-23.4 4zM41.7 17c-1.9 0-3.5 1.7-3.5 3.8 0 2.1 1.6 3.8 3.5 3.8s3.5-1.7 3.5-3.8c0-2.1-1.5-3.8-3.5-3.8zm1.6 5.7c-.5 0-.8-.4-.8-1 0-.5.4-1 .8-1 .5 0 .8.4.8 1 0 .5-.3 1-.8 1zM71.1 16.1c-1.9 0-3.4 1.7-3.4 3.8 0 2.1 1.5 3.8 3.4 3.8s3.4-1.7 3.4-3.8c0-2.1-1.5-3.8-3.4-3.8zm1.6 5.6c-.4 0-.8-.4-.8-1 0-.5.4-1 .8-1s.8.4.8 1-.4 1-.8 1z"/></g></svg>`,
	}
}
