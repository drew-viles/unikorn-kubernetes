// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+x963Iau7Lwq6jm+6rWObUBA4ZcXHV+EONksZfBToyT5WxSKTEjQDAjzZE0xuOU3/2U",
	"LnOfgcH2umRvl/9gkFrdrVar1epu/bBs6vmUICK4dfLD8iGDHhKIqf9sN+ACsQn00GX0g/zeQdxm2BeY",
	"EuvEmq4QMC0BgR5qgXHABZgjAMEtdLEDhpMrYFMiICaYLAElbghcukUM2JAjYK8gg7YctDEjJPDmiHFA",
	"GViF/goR3gBcQCYAJA5AxAFbLFYAJr1kU92rodrIgQXwKBcz8uo4BR1gAlxElmLVshoWlrj7UKyshiXR",
	"tk7S9FoNi6H/DTBDjnUiWIAaFrdXyIOS/v/P0MI6sf7fUcK8I/0rP9oEc8QIEohn2fbw0LAkDxh1L11I",
	"UB2m6ubAl+0VaxsAL4Ao/ORQxAGhAqA7zEVDtiAAC+DBEMzRjGDPd7GNhRsCmyEokNMAC8oAuoOe78p5",
	"iuYP86gFgEuICRfyx/RgMyJWUOSG/ImnPDclf8i8+4yukS1qTLlpqddROcYpYH8IsgwtMSU1cNUNd6Ga",
	"gPoDMH3QIBEX76iDkVZYSnZPU3P6STdRP1IiEFEfoS8XBJTEHK25pOiHZRZD7ud3AXH0l1mBaarF0Oy0",
	"2q221bBuEeOaM51Wp9W2HmIeOGgBA1fIb+pRnJZITWaW9aeZpW9YABLF3SrwWmoexZjfYmae6vX+7NxJ",
	"pqtpVEopi9qaRRlST35YCxfeUq1+T6xlq9viAhIHMkeKlweXyPyE7E2ze9x+3ek1e3O0eAPnHUW0wotb",
	"J8fp0W47re7rVleOt0BQBEyLCgwE5TZ0MVnGXMpuA1KOkdhStlHyT5QEcsRu1e74L+tNS/1ZDfWp1+pZ",
	"3xoWoQ66ZGiB7yShb7utzqs3ktyjziurYfnUSX5st9TfkYQgwWI71fO17Kk7KtSpjwgX0N7oufL8QKDB",
	"LcQunGMXi/ArlSy0CL2FVsNCdwIxAt2Jxn80lFS9dTrH7bndPG53nGavb7ebb4+7b5rw1dtXPbh41e+/",
	"fiunibqBVwn6oWFJgC6FziWlruRDjpU/LA/eYS/wPqWnw8Mk+137oWF50F5hPfMO5ooyju+RddKXv+aE",
	"odda4eXKQ14LdtrtVmfZ6rSX82cSjPxa/fZwuH4yS6psySbrLt5oD1q3l1rlP2q1GtK8sGk2jvqaKO5Q",
	"pMhgVI8Mpam5Twk3S8+2kS+Q88l8WbW9aNAryMEcIQKibmrH32LXldv+InAX2HXltzwk9opRQgPuhq0Z",
	"uaGBsoB86rrKaGKI04DZSAHwKMGCMoAFl8aGCLgyiSQbXCTRaElGFXRdGue6k/CvQ9XjsVSP3+pOUwFH",
	"XjZjA+BiLgBdgFR7MNcd8rQ+ksqCjNxiB3EACYCuUkgC38pp0FCQA7igDC6V2SObMqBtS8wFw/NArpeo",
	"BbQZ5VwanwgUl1MLgPdGtwOpJpow0l8ibABMbIY8RAR0ASfQ5ysquLYbob0JfGmDOphDszBteotYqA1L",
	"voIMOWCBXQQ8GhDBwX8xBJ2jLcMCAQ+S8L+lxDvUDtQIhvaVED4/OTpyKVmuKCMtTI+shrUKPEg+IejA",
	"uRvprHPTRKoyWzPu10n3a/jO/zps4+mH9/2vv/9zMb4aLb9+eN++ueoEN1867uXVP8c3v7uujQd3I/yu",
	"N/9yF9j3bQx//dS2h/T2/Ng5dsL+8Tjs39qefTteD7bj07f3jmfj0a9f/a+/O6fz4+Xb0XqwHJ8O7i6m",
	"H4Px+ro7nm6W4+l1/3w96F1Mz8LRuvfG+eC25x+u/wG/TG7n6+1t9P/lr+9Wzofl8qvn8vmwjUf3n73x",
	"etS+kbhK3Keb4/P1WXgxPOMXw0EwWY+6F1/O7sanve14uOHj6SAYDwf98+GAj0+3d+fTs+Biet07v+rd",
	"XUzH9xNvKyZXvfBiOO5PTtt35+tBZzLc3J8PPwaT6cfeZLrh47UdXEyX9+Pp59XFVa8/Xn8ML662/fP1",
	"JpwMRwns097deL3pXcjP65vtZPixD4fXwXg66t5MN8HFdNOfhKpf/2Jqyz7b8+EZP1+fdcf3g57EbXK/",
	"OR7ff+WTq972Yrq8m1y1w0nY64+HN+1xe9u/kN8Pb+7Oh8vt+frj/fj+uv1xerY9Xw+2F8NNeD5MfzZ4",
	"DUt49Jni8/veG/vD+zY8fefBL3f88mq0nny5CcfrT6sRfre5vPrnZDy178/XN/3J9IaPz5bh+LTXmawH",
	"x+PrM/m5O16fbSdX2/TnrRl3ez4cbc/lfA9vjj+vz+4vTnud8XrZnnxJ9cXb9OeobzROdxKmPreXd5P7",
	"cTBZbzoTL4bBx2tF011x3OvO+TSNQ/L5o/r+JhwnuJu+A56h+b0vxmGvPZle88nwLJhMl3fn01EwmQ4k",
	"r49vDO/Hw5tI1hI6rtrH5+vN/WR63T4fLoPx/fV2Ml2NpTycrwftyfRj53xod6TMjb+MhYQzCXvbyXBw",
	"PL5qS1i9iVwzw+XdeHgjf7+bYCljZ8eT7lZMcO9+omm4n5z2epPpoHNxpviyHa9vOpoPg3Cyvo5l7WK6",
	"kfyTON6N18vgYnrTHa8/0/NpJKemz3R5HPdXn+P1I+X3+GJ4HerPg87F8P14omB9bE/ur/nkXsLaHE+m",
	"K34+/Xh3vv64HU9vwvPpMhivb7ofd/Jse3dx1euOh3bn4mrbkTJzMXzPY55P0zw/u4/kXX+O5F3iZfcm",
	"92dqrqSOGU/f8/FVT+In4Wr9sN7cT1NrYyLlaDjqT9YTPpkug8n9dX9yfyPGal2O7ybDjykY7RjGx/34",
	"HE/C3p2cnwnetsdXiiY4wm/+can15T9Ol//zP1bDcrGN1J5oDXxor1Cz22qDc/NlbENGGr/ZafVbnWYn",
	"2dq1sZze5/utjrQ1H7PT79vj9f7novRur7f5OXSMIfmYXf6HhRijzDqxMFGenu/GTpOnDvnL9yxKkRU3",
	"p04ITJf6FiiFgVh1z9SIJfR+SgNfQCzNQN1Ve6EUDQ1AmTH/dOvYdWX8UzMCYwNR261ggZHraHbZlCxc",
	"bD+RWRGUCi7BxDZVri6JDIeedvoB6EqTI9SuNv6M3DNDRshx42cjVKwQa4CAB9B1QyBWmAMPQcIlYiFY",
	"wVuURbGVP80/jlvP4ncpABkEgl77SwYdCa942mtY2v6PHUiYkilWTbrt7nGz/bp53Jl22ie9/kmv+9Xa",
	"AUDbvBIj5Chn2vO4fAZZF2yB2/yR9vrT+N3+2/H722MYvkeTZjivVcKCsjl2HESephNiMBVKIeCIAZsh",
	"BxGBocuBQ5XaipdfrK58hm+xi5aIP7tq3UIOHEQwcsA8BLIPZZgbxaoVg7pkADYMuG4kUcs0nBFBN4hE",
	"yGOyzKLPbeojdfqGBAwuR7HGVhyQ6pr8kpA9IwTZiHPIwhThgBLVJT5J+i4UC8o8NWOYaH/YlfLeKaKf",
	"NnfaDfhd/1s+fWY/EhRo6m0XYu/Z5mdAQEDQnY9seY5W4wNq2wFjyMlODMy0FAwSjhERpg8kzozIljyw",
	"bYQcyUe5GwkWtsBooSFhNQHq3gly1AC+iyBXB3nKBMACQHXIx5wHWjkVvGLyyC63GbyswXdqCySaXDAE",
	"PalBSlZniUdNgw8YjJ03m6K7+7l3pZ/Y392wbMTEGBK4RMw6WUCXo4a1wC660j6X+DtMlgxxHv+fED2E",
	"fDWnEuHoN3KLHQwvfMSgoAlYn1EPiRUKIigv3vZ63vYDtsr+V6uEqeVb5Ysb/xFu/DK1U65o/giD7EXV",
	"vKiaF1Xz91U13x6ta/acQIpKRx9DCBXvaUCcp1myhIrvCwmmwoxNeU6Qk7gpsvE/z2bWXhPltBIULDBx",
	"QOIXaWXWyjuX2hujO/ISzZ92RadXQ+3JjFEqoLF7Ui98RK5kPwBTHcE9jQ6ZMeDTcp3wLGQ24v8vL4bN",
	"Tv6L7t+KEWdZ3fdYBmCnvs40vBip4yMSj2FHHuu63Ig0PTBbVY4Z75WueywPbF/q6V7DaNFuu2Et1Ved",
	"hubPW/jGfnX8ut3stV/1mz2nB5tvHdhuvn71+o2z6LVt561UGB7yKAutk+NuzKtKvfsI3hki67JM6/8c",
	"o0ZS2T+aTzoWMt4Du81ud9rpnrR7J51juQcqZsFXvcXb7qu3zeNXqN3sHXe6zfkbp9Psd523x07/1dv5",
	"a7nteNTBC1wCrdM/6bxJ7ajBPOh2272m3G76rVfNpR80+91+602/1e43X9vI6XX6vczFwo+UpWQ2qn5L",
	"GibaSBoyfCstLysGc4i3LMfLutOhtlmur/QhQyq0Agos9btxcmOeddPEA/2GwkuI2RN1nBc2OV81Nyh8",
	"jPBFONQld4NC4MsOihQTO/N0CuIgnG8Hhu3sQTxqpbDVgZpPRDbYNLdI2QLprxDkhyBvMNmNu2mkUA+I",
	"8TXeoydaQtC2Eefftbez6qImECtERBQ+o318z+fUK4MeOUM1esaZuIIcoDsfM+S0UhcNPEduMbJqQNI3",
	"gi1LSaqPmDARtJnW+c6fEZtTjkDqWzkdW7nCFYqpyCLjiFXxYCL0pTRwwaTB/1AIl8mPM0z/DFxMNso7",
	"nBtCQpaqAwopaQyXDVQScJMf7FfZBDDTJkNDFNpcAKsDdQq8BXPI0aseQMSmDnLA1ecPQDZtATCVyo6v",
	"aOA6QG7nABMwp2IFXLxc6UB3B7KNpNHTTItJm4cClSER30eXRc+ZH0FAHMTAdoXtVWGKMAcMKVeuU0ol",
	"KeXXNcH/G9Tkk4BLfsCl9lQ2f8juazW7fo66RHHhOv7wX5qIMkHILr68TCbsNbOdwupbTCmdRxGVpZ6b",
	"gnioX3IxeLylhUPOt7pglVKEudZz5rQVDd0CJsIPeJBtkDMjkAOfoVuMtpF0yUPZHAGOXO3xn4fAnFIb",
	"cT4FXQAXL1AU/pftOiPqJkdQAG8pdkCQunAJ9L0eV3cDSB3WnAaQx3kPCmzHv+v4THUhAfBiRiAgaItY",
	"RIhiQcQOfa2t7QSsD5WYRFS1wJcVInHjX7jBf0YUAVoH8kbMKjOyEvslBVCyFdnIiTCTLZeQSaq51l1I",
	"rBCbkQINEhdDob6bSqaDMollUXki4lwszvGiZPIVFWpyNdEKuNOki6akI7PeHShQU2AP1V+T03R0a+Vq",
	"NCwtRU4yIYdfahISaHNKXQRJapWWY2PAmDYl6JQv0whmrSWWvWguYzj3ka1s7oaZSq5kNZ7iimhYMHDd",
	"vETJdRHLiAqLMECcOCcrYAwR4YaptZee7GjVlWy6MOQXiy8IbfZqvITkYdJJMnM/u3iZStoZFNywsEAe",
	"PzgA2UrwgYzBMIfOEPmIOIjYuBwnjvIoGe0QBX1jaaGroO85kutbHyVypsGhqMdYhXXRD/eZV8CJmxYn",
	"vXol19hay1bPnjXzCdnU8xBxdvGcRY2Qk0FDsd9cKCfchwuhtMafyPypsSgq8JcGh9LWC+wKJHmVC4BL",
	"obZz5gRclls01ah9rtKHOdApnZi3vbPr4lDeYR3/wTITXRNISjr2qfZUr184+BW5nkq4FPWVfU0t/zll",
	"Bu7XEYmR9Aj5i+Zu9wzv06ClYlYTg9KhS7V98bgEQ2XLyE1oi9BGmXnSgABbTBy6NdrTR8zDAlB1YaaV",
	"KpXr2UdM2h5yIysRygXDDgz3ESJH+6IGk3h7lBzch0MRsMN7BYePJFYB44f3CtDhnbbIIQd3K9vM89e5",
	"e+ID85NYejo5eEvfZ4AdBDDdd6ddq0JgjXQnt0Elqjm5v6x3+xbFv17pfhUWaZET3/bMz04tkQ8nrKko",
	"shGiRU2xogEr3XjkDxH3HBhKo/56eirHNZeu1kk3uW21TtoxbEwEWiJWGt9QHKosIMvcIueSA/OSiQ++",
	"Lx1cjiqP3H8roc6v2lp+57G+476k1M0HVhzEpSghrnp1naaLdpQeNpMoiYOGNnddhfCGg4DEnvXnWNuF",
	"IIQDkfmS6V1bVaTpT9iZE4w8bmX6pXQN7FqEg8uRtJEEJsuyVee6dItMyEqZwroy5pXjMMSVi0k1VE4h",
	"2TfxRoNcRubgcrTbyh5d3vbA6Wj4KQe9VAI9TEYaUqeo9Hig+DNIcktVIE8lNYSSZhT2BH5v9dtvwdVg",
	"oolynIgWyTlbsmqhklR3ExNDORT7h1qTnA2TqREEG0kS8Cl1QSrMJhceC6K1n2oyI17ABYAuV2Zh5PvC",
	"2m8VjRDpo6JQFaJ4ynZz08hUb9FHTt0+lq2kmMxSXb+qrGCNhNmmUsxO7VKFiKHS8XWjeuMrJ2EyuEa9",
	"bPCcPshj0ijwptYaf59S/RXnH6q+ga4SYEri2YnvS+IwxYIK2CVbZzoMR7ZpmkblPsBMXF8FFNmm6elG",
	"5VAykYAVUC4vrka/60TsOeTIkccWjrlARMRJ4v8V5VL/d/k4cXRhFb0EmCaRneZWoVwamFgBNqchnahD",
	"C4BPWmp4PK48uqWYao5oZi2Wo5KPg8xjMdL+GoXG5PNoOBqAuHEZvHQAZdVkxE3KUKql2yapCMycbG+K",
	"es1snTu2tHwYZ7UJPpxc6UOFbitZHPCdej7uYnqoLczsXrW8ROng0Tx0wwizC0ps/Ch8Esh+UjaK9bXU",
	"Bjp3qb0pd/Un8aiHjOdT51HD5aJcDxnSdH3EsHkjLOFxHqE0Pxp5SamlihNj9KCTTxIdsusMVBn1WzDY",
	"dcNipJwpXZI6Wuq7vqw5YK4BS2ewJLA4P3p1bBoYDcvFgq9+Q2H5tXsC7erqV/AbCqVYGNeySkJyXWDC",
	"fcvXWFU8cyFoQbV7fp7lxK9qEisRLeN5LVlMn0WqzcJKq3CvKXDYWTzVV56m907Hl6yFmp+VFri4RYyp",
	"ujFps3OX7LpwjvSZDjoO1gbRZXVUi7SeVF+pZILyk+8unDco1D2BHlgZjr7vhsBo7Hj9p0An85mKYX+M",
	"P6D8KJ/F8IBLowSfg2Vv5ya772xS3/W1W/73nRFLgu0PwvoJeJYZAlXlDHef7p5aVRU8Z4VNsKPApgfv",
	"ztU/1skr7VWM/u2ULLJKd81ubsSBDtopVLKfZtJXyk6B6nIkUxJsC+PKp/WDMRykS4UdOpDqd8hAvIIv",
	"0yTsICpilveUg5HKHnd1XrZpFF2kz6xrsiF0S2YWCIjAri6wFSOrIsRsSmwViGhyuxly0S0kIn1GAiOh",
	"q84qyDrJWN+j0hmZJQlBmCxnVqaAmy5tIffigCOwNZVr7RUkS+1vmqXTiWZWC0xjOmZEQVEei8yYCs/C",
	"sIZ4J9D3wgQEvuS6gjgjadbo4fXoQ+RnwWx1QJKexLiqCOZgjiRcn1EbcY6c1oyM9KW5QjANU0WAziyA",
	"F7n07Gwid4JqmFzbtWZEdY8TvOOU7hpGSnplxGJVpvnTkaoFsfuACGLYNth6iHO4LLl1QuW9B0AqDmR6",
	"m60Y3fmQOFqPqdn7dTq9NE1s6qAWMERDFjkATMOLgcQ0CvWxTWmCeaCDzjVcZBScxI9hJCALI2mw1RFL",
	"yt/gcsQBNYFQyudDOUpCiKT467EkpYgEnuRosQpOOiD5u+3KibEaheDigPDA9ykTSPbVYcvf1Sw0Ypiq",
	"LoI5zKSKDQjk+ZRBht3we0DiSj+pjvGo0RdLBonIjaq+i4ZMZ4GlatXIQz51vstfjcs4B8RDDoYRkKSg",
	"xbeyE0YxnLoqvthIlIkznkcFIxSE/UJeXZShVNCr8pZK/b07spUOivApHEaeGOezI/tqh7mzO/eqpt1T",
	"zcAS+6cqLWoPs/MHziKvsfMcB1ZSdVTl5WDqzRp2otyenVNXyBSrNXMleWKHTlx+LnbNm07J2jNdOhGr",
	"xDDzq4yYxB1/ennNy/37UcZvyXJSpTllb+SvkIcYdIFsLY3TD+/KoS1r4PLh8pqrkvuECuVxl5sDUvsK",
	"MSu+CLhMEiXYQMfNa95UCWCUULebSt1KUYcryKvWPQaBQ0W3oWcvRtHMx06JjvL3aglynL13qPgakdwl",
	"tSpxrUxonXy+WsVpApVNqoqwiuK6Ve/MMQLodLnoaKUD4GPXoDoZ6Lj7GZkjsIC3NJA2H72Vwuc6iEUZ",
	"dNDUXdOmqT7xGdtVXxgsdGx94BLEtALGuRySnaeLPRKrKasS2DipsS57XMjlQVJ3e47Adw26MoihbjJJ",
	"Vlgq80kyK8KIRooNe9JEKgYp3tkomqJsCQ8J6EABi8KZzvos403qKoUjDxKB7ThMM5uUNFflthzgYIZs",
	"4YZ66rT1G6rb4bRzJBPFVHSbKrdA8qpIxjYv33gzaaqlWlm1AI5qUj/UNMWg3Cj7p6em4jowz/VQ9aZ1",
	"1y7tZjJV92zKV1e/xjmqh9iqUZ9nM1HjxNpa3E2l1R7KuYgvu3iX9q7WC7Ew/tKSG0hjn9TCTd9KWLmq",
	"JdUWbG7X3uNZS1U4qQaZ1Zt7ILLKqIpJbC6VXNekLJLKmPIUj2PVlNNMDpJr1gELRr3UdqKro3KzAo03",
	"aI6Az5BURdGp32yZiHGTyBXpI1Ay9D5W5MSdJTEeEYFp9memd+eqMAKxew3rq6SSG5xH34I95r5Fl9wp",
	"RKFLg1v+tMMuzTFPASrjSpT0XrIizU/VGuxPiQvegfROvZZk3ddUZqlHO/I6TGfBl42mfzlEyefecnqk",
	"io+S93cwIE7dr0m/IbKE/FTIfJmLMUl/MLljDPrSLDHOYYLuhA6IXsSFSEsTOvchqCKvtZQxUa9xfhWo",
	"ng01WJGrKrrCDhgW4ZWtksPlINo9m60XUKrrtc84IpFH3tI5gkyFvmwQAdmiBsokd+k2uhpWrkz1yynV",
	"IdGZL6+Zm3rvAqtitSJsBQRvKCNN26WB06JsaWofHN12jzL9rYal/Jxc1wpWGD0CpuqXCT1SP+kqC5gs",
	"aLncpzaAKx0youI/o+IFSXSKtu5cyReQvg9Qtq9KRbZD20UzogPdPEQqLxTVBYYcBXPg0qXJJlXSotza",
	"i9yEzEiERSM2I5M6t/FeJsGoC5MlEnKPixL6zFWOZIshw4ZEHTnVhU1UtRfOuWDQFmUsSVJhBTXHW/2G",
	"jqI11WNGEipNBWEOVFSpOUgY1/74HJj6HAqvGVkh6JhQOyxclLUKUjOTKXbYbnV1sUO5iapEAuu41W4d",
	"y2UMxUoJ8BH08dFtJx2qb9Jbj+yqXIY4NraQExvLg8R0iUo2qHOVy7506Ry6ZUm1OuokZpI2dDBPcu59",
	"hrjkCwT2ikqq6SLObFOd09LX0hHmepGPHHU1IwY+/twZFOg9jev0ZR5j6rbbVVorbndU/QzSQ8Pq1YFQ",
	"8uqA6trZ37W0wstDw+rXGXdX5ee0blWFbMq16r++qUKid81MhYjmktHAt04sD2K1P+2StJ3ZW9kH/f44",
	"octmHv2popdNd3iRvz9J/ng93VZTvjL5+MkDIzpCOUkKpyQTO7NXRvhTJeJFFiplwax5vyItceT5rjJV",
	"ch49kDJiYuEAp1mnn9n8VUX+xACqMHoaQKwYDZarjA5pmBoa6qOgIIqnac1IfjBpQzC0QAwRW4VAab2U",
	"K1NgFKZ+RQ4tMNGn2hnhdCG2kCWVdAq5mCCZMh3xVFKPJ67/MSNRBZFFQGwdcSgtVAA+GRxNhAm60y8v",
	"l7zlrNz5JW8z6zJNnFMbqzfyUgfeHYsz55HdYrEyiTvKus6AKV+SpxlZecyaLH+S5C9clb12b3/nQhnd",
	"P385+5SXTOypEgeur45yAnTA/F5SXjXB0evJYTWxqQeWj6pfV34oiEu3hgrPPwP6l0rK21rynX0K6idR",
	"/Ec/8q+cP2hpc5Eoc+Sp73n+4XftfK0WvBkxBqXsCLkNdTRX7IKIg/5M4BxMLjORs8Nk0OgUJfi07OX2",
	"n1sGfw5t9WJD/NvZEB+QOHi91zMk9q/SAw2Ll5V6gF0RJ52ptmXDJk2O8ptEkp3w8K1h+UGJ3FxHQShP",
	"ME+CulLzYq38bPL3XNZK5Krd4c+o5cPQNoeBlRHSqGBjoY7SI1RcnGj0GFW34ymhF4H7cxVenUNZlNp2",
	"uCiVH8t2ytKjFOBvxUfw/m204PH+zsW3Ov8DT3ux/jz6YT7VPASmYvjS1iA8SNLrHuAiWT9NUHw50/0l",
	"Z7ram+oHddVcJiJ/2K66Uzoes8G+iMlfsL829ndO5rn2QSQli4/ZkYMniOHL3vxyQinusErb6MprTzi6",
	"ZLTsLxxkA/HTjx4/l2ZNnmt+Hh1b8vzzi3D+FNq27gKhbAlJFNJX53KBgHQX5S2Vp3MseCrcLJPonXkA",
	"Sb2mvvu24CKN0otC/SsMyUFuliN/OgTzgMtPHOjYThOYqK+DeBK+GMVLzwgSdku3miOXkqWu8gBFZoAk",
	"qkkHs1PIOOICLBnExAw1IwvKwKd3A1UFuc4pP0PC4SIZn/T/zQTyJzhDS/1FaHOu7DLBAvSQUlnV8fqp",
	"Oy3TqPw269JAaADIU8+/aCl3ZsQUTPktnYSn+jZUoljq2sZ4yU1welqsfuEzkhHxvQ7OiK64Bmi+e7mN",
	"EBHzqB2/8NzhTxZQVUcNRKIQh1RH/A24Mst2sznWAhk+P+rYYCC8BF/8TQ8L0WI4+mE+1Q65kNIUiVnq",
	"+takkaJ0YgByIrnTumMWMZeb/HKpPFrg4JiMGUkFZRTTfGvFaUQifpmQ/2KA/S0PDSkBTdv9KWGuzOn6",
	"pH8wu506oNLAiSwwU5tqieiSQX+FbegCylQlK5PzAnzIBI4qWmFVd2kLQ50ujaO1gD0s8C2KXweMnyaM",
	"VkecHA94YK8A5DOSGdSlNnRVaU+joHn0chHTJWAdMHfpXMdy6PKq0tSckTNor6K8uBXk6mRCtyTJvMkv",
	"UCwaakswT8smaZsqFmRGIgDGtLUZUolN0OWAy30fkyUHNiTq9cbYHkhlq3LXBOXDGeEryFTYFA2cOC5m",
	"u4IC3SIGPGSvJKmeWvpxqHb0kJfuFRGy36JIJe2V2w5GFh5lOuTfHv6pLIfiQjn6oT9o/08637ap6iod",
	"qXLGTV5V7vyd/DmuaF4sFFU3WN91TbZwSW2nir1EivmMxMbn3un+FBNaqDqlqIhKuj9GKuK86DSkwjD/",
	"gaH/h+n3RBbruXUindBM3i15hIgbRXpIAesDxNquArJLrsGziLXB/2kSfVpeKfpFmP8GwhxVU2uSyiJs",
	"Z/kicAeIbqFY23NLbKGG3JMkNQ/tRUL/egldVNVUi3SrqRTzCIVqQGsbsY5g1pfLqBLck8TRAHmRwr9e",
	"CnFFfaxICHV9nkfIYLqY1jPrRlPT60kiqGG8SOBfL4EbFDb98jJiv0UVyw6Qv7jM2HNLXVzt7ElyF0F5",
	"kbw/QfKqXL0agZIK/JTpgi6mPox+LC5dBka1iK+1VRqUjXwRRHX5TVHZ5LYzLiaYuiOKU5d0vZbER7RB",
	"BGDOg1TB+GxFHwC+qHCoGYGJe1mOmap1my8PMTLP/lEi0J2InhUoVmVpAKhdxkn2h3I3pxjgwlDXVBdJ",
	"YR0vcAVuCkQgEQBz6pq67ZA4ZfVkijV3oncNk9u4kpSxiKvR8wEmLEzyIRsMBtJPBaQvnrXTO5+1pnyb",
	"jjy3RslZbggoS2diNcCKbpUzTjk0XShM2QVGob1StX+kLlm46E6VitRFXsuq7+gcrxW8VbUW7BVVZemp",
	"h6Iibfo1Gn3fF9IgGRmnGA7BAurKD+o1D4mNqvzBJAmIYSlYqiqXE91oRDVXrVMj3yqQpCj+qUeuCtWJ",
	"MsWUkjdDs2WVVG2wW8gwDfiMxEBiv2eq/FC0LKL6WdGbFNESzHpgbzGTa2xGTGVCUwRJckDbGS3wZYVd",
	"JLmmPK8eJHpNRuWK4rs+9fpbfBEjJT4aEIvomQz1fr5ckRJLCXKBGVdRDVzOknJol3GIAymSlDn6MTpV",
	"v4mAwFfpjFCgVPG0AiMS566p8cIDz49uINVclmi4eGaTqbuMELtMIWY9fHv4vwAAAP//WbEWQiK8AAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
