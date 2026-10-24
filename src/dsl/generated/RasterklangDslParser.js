// Generated from RasterklangDsl.g4 by ANTLR 4.13.2
// jshint ignore: start
import antlr4 from 'antlr4';
import RasterklangDslListener from './RasterklangDslListener.js';
import RasterklangDslVisitor from './RasterklangDslVisitor.js';

const serializedATN = [4,1,37,469,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,
4,2,5,7,5,2,6,7,6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,
2,13,7,13,2,14,7,14,2,15,7,15,2,16,7,16,2,17,7,17,2,18,7,18,2,19,7,19,2,
20,7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,24,2,25,7,25,2,26,7,26,2,27,
7,27,2,28,7,28,2,29,7,29,2,30,7,30,2,31,7,31,2,32,7,32,2,33,7,33,2,34,7,
34,2,35,7,35,2,36,7,36,2,37,7,37,2,38,7,38,2,39,7,39,2,40,7,40,2,41,7,41,
2,42,7,42,1,0,1,0,1,0,5,0,90,8,0,10,0,12,0,93,9,0,1,0,1,0,1,1,1,1,1,1,1,
1,1,1,1,1,1,1,1,1,1,1,3,1,106,8,1,1,2,1,2,5,2,110,8,2,10,2,12,2,113,9,2,
1,3,1,3,1,4,1,4,1,5,1,5,1,6,1,6,1,6,5,6,124,8,6,10,6,12,6,127,9,6,1,7,1,
7,3,7,131,8,7,1,8,1,8,1,8,1,8,3,8,137,8,8,1,9,1,9,1,9,5,9,142,8,9,10,9,12,
9,145,9,9,1,9,1,9,1,10,1,10,1,10,1,11,1,11,1,11,3,11,155,8,11,1,12,1,12,
1,12,3,12,160,8,12,1,12,3,12,163,8,12,1,12,1,12,1,12,5,12,168,8,12,10,12,
12,12,171,9,12,1,12,1,12,1,12,1,12,1,12,3,12,178,8,12,1,13,1,13,1,13,1,13,
1,13,1,14,4,14,186,8,14,11,14,12,14,187,1,15,1,15,1,15,4,15,193,8,15,11,
15,12,15,194,1,15,1,15,3,15,199,8,15,1,15,1,15,3,15,203,8,15,1,16,1,16,1,
16,1,16,1,16,1,17,1,17,1,17,1,17,5,17,214,8,17,10,17,12,17,217,9,17,1,17,
1,17,1,18,1,18,1,18,1,18,1,18,1,18,1,18,4,18,228,8,18,11,18,12,18,229,1,
18,1,18,1,18,3,18,235,8,18,1,19,1,19,3,19,239,8,19,1,19,1,19,1,20,4,20,244,
8,20,11,20,12,20,245,1,21,1,21,1,21,1,21,1,21,1,21,5,21,254,8,21,10,21,12,
21,257,9,21,1,21,1,21,3,21,261,8,21,1,21,1,21,1,21,3,21,266,8,21,3,21,268,
8,21,1,22,1,22,1,22,1,22,3,22,274,8,22,1,23,1,23,1,23,1,23,1,23,1,23,1,23,
3,23,283,8,23,1,23,1,23,1,23,3,23,288,8,23,1,23,3,23,291,8,23,1,24,1,24,
1,24,3,24,296,8,24,1,24,1,24,1,25,1,25,3,25,302,8,25,1,25,1,25,1,25,1,25,
1,25,1,25,5,25,310,8,25,10,25,12,25,313,9,25,1,25,3,25,316,8,25,1,26,3,26,
319,8,26,1,26,5,26,322,8,26,10,26,12,26,325,9,26,1,26,3,26,328,8,26,1,26,
5,26,331,8,26,10,26,12,26,334,9,26,1,27,1,27,1,27,1,27,3,27,340,8,27,1,28,
1,28,1,28,5,28,345,8,28,10,28,12,28,348,9,28,1,29,1,29,1,29,1,29,3,29,354,
8,29,1,29,1,29,1,30,1,30,1,30,1,30,3,30,362,8,30,1,30,1,30,1,30,3,30,367,
8,30,3,30,369,8,30,1,31,1,31,1,31,1,32,1,32,3,32,376,8,32,1,32,3,32,379,
8,32,1,33,1,33,3,33,383,8,33,1,34,1,34,1,34,1,34,1,34,3,34,390,8,34,1,35,
1,35,1,35,5,35,395,8,35,10,35,12,35,398,9,35,1,35,1,35,1,36,1,36,1,36,1,
36,1,37,1,37,1,37,1,37,1,37,1,37,5,37,412,8,37,10,37,12,37,415,9,37,1,37,
1,37,1,37,5,37,420,8,37,10,37,12,37,423,9,37,1,37,3,37,426,8,37,1,38,1,38,
1,38,1,38,1,38,5,38,433,8,38,10,38,12,38,436,9,38,1,38,3,38,439,8,38,1,39,
1,39,1,39,1,39,1,39,3,39,446,8,39,1,40,1,40,1,40,3,40,451,8,40,1,41,1,41,
1,41,1,41,3,41,457,8,41,1,42,1,42,3,42,461,8,42,1,42,5,42,464,8,42,10,42,
12,42,467,9,42,1,42,0,0,43,0,2,4,6,8,10,12,14,16,18,20,22,24,26,28,30,32,
34,36,38,40,42,44,46,48,50,52,54,56,58,60,62,64,66,68,70,72,74,76,78,80,
82,84,0,5,1,0,35,35,2,0,1,9,16,16,1,0,13,14,1,0,1,2,1,0,23,24,498,0,91,1,
0,0,0,2,105,1,0,0,0,4,107,1,0,0,0,6,114,1,0,0,0,8,116,1,0,0,0,10,118,1,0,
0,0,12,120,1,0,0,0,14,130,1,0,0,0,16,132,1,0,0,0,18,138,1,0,0,0,20,148,1,
0,0,0,22,151,1,0,0,0,24,177,1,0,0,0,26,179,1,0,0,0,28,185,1,0,0,0,30,198,
1,0,0,0,32,204,1,0,0,0,34,209,1,0,0,0,36,234,1,0,0,0,38,236,1,0,0,0,40,243,
1,0,0,0,42,267,1,0,0,0,44,269,1,0,0,0,46,275,1,0,0,0,48,292,1,0,0,0,50,315,
1,0,0,0,52,323,1,0,0,0,54,335,1,0,0,0,56,341,1,0,0,0,58,353,1,0,0,0,60,368,
1,0,0,0,62,370,1,0,0,0,64,373,1,0,0,0,66,382,1,0,0,0,68,389,1,0,0,0,70,391,
1,0,0,0,72,401,1,0,0,0,74,425,1,0,0,0,76,427,1,0,0,0,78,445,1,0,0,0,80,450,
1,0,0,0,82,456,1,0,0,0,84,458,1,0,0,0,86,87,3,2,1,0,87,88,5,35,0,0,88,90,
1,0,0,0,89,86,1,0,0,0,90,93,1,0,0,0,91,89,1,0,0,0,91,92,1,0,0,0,92,94,1,
0,0,0,93,91,1,0,0,0,94,95,5,0,0,1,95,1,1,0,0,0,96,106,3,12,6,0,97,106,3,
18,9,0,98,106,3,20,10,0,99,106,3,48,24,0,100,106,3,24,12,0,101,106,3,32,
16,0,102,106,3,34,17,0,103,106,3,70,35,0,104,106,3,4,2,0,105,96,1,0,0,0,
105,97,1,0,0,0,105,98,1,0,0,0,105,99,1,0,0,0,105,100,1,0,0,0,105,101,1,0,
0,0,105,102,1,0,0,0,105,103,1,0,0,0,105,104,1,0,0,0,106,3,1,0,0,0,107,111,
5,16,0,0,108,110,3,6,3,0,109,108,1,0,0,0,110,113,1,0,0,0,111,109,1,0,0,0,
111,112,1,0,0,0,112,5,1,0,0,0,113,111,1,0,0,0,114,115,8,0,0,0,115,7,1,0,
0,0,116,117,7,1,0,0,117,9,1,0,0,0,118,119,7,2,0,0,119,11,1,0,0,0,120,121,
7,3,0,0,121,125,3,14,7,0,122,124,3,16,8,0,123,122,1,0,0,0,124,127,1,0,0,
0,125,123,1,0,0,0,125,126,1,0,0,0,126,13,1,0,0,0,127,125,1,0,0,0,128,131,
5,12,0,0,129,131,3,10,5,0,130,128,1,0,0,0,130,129,1,0,0,0,131,15,1,0,0,0,
132,133,3,8,4,0,133,136,5,12,0,0,134,135,5,24,0,0,135,137,5,12,0,0,136,134,
1,0,0,0,136,137,1,0,0,0,137,17,1,0,0,0,138,139,5,8,0,0,139,143,5,28,0,0,
140,142,3,22,11,0,141,140,1,0,0,0,142,145,1,0,0,0,143,141,1,0,0,0,143,144,
1,0,0,0,144,146,1,0,0,0,145,143,1,0,0,0,146,147,5,29,0,0,147,19,1,0,0,0,
148,149,5,9,0,0,149,150,5,12,0,0,150,21,1,0,0,0,151,154,3,8,4,0,152,155,
5,12,0,0,153,155,3,8,4,0,154,152,1,0,0,0,154,153,1,0,0,0,155,23,1,0,0,0,
156,157,5,4,0,0,157,159,3,8,4,0,158,160,5,12,0,0,159,158,1,0,0,0,159,160,
1,0,0,0,160,162,1,0,0,0,161,163,3,26,13,0,162,161,1,0,0,0,162,163,1,0,0,
0,163,164,1,0,0,0,164,165,5,32,0,0,165,169,3,10,5,0,166,168,3,62,31,0,167,
166,1,0,0,0,168,171,1,0,0,0,169,167,1,0,0,0,169,170,1,0,0,0,170,178,1,0,
0,0,171,169,1,0,0,0,172,173,5,4,0,0,173,174,3,8,4,0,174,175,5,32,0,0,175,
176,3,28,14,0,176,178,1,0,0,0,177,156,1,0,0,0,177,172,1,0,0,0,178,25,1,0,
0,0,179,180,3,8,4,0,180,181,5,12,0,0,181,182,5,24,0,0,182,183,5,12,0,0,183,
27,1,0,0,0,184,186,3,30,15,0,185,184,1,0,0,0,186,187,1,0,0,0,187,185,1,0,
0,0,187,188,1,0,0,0,188,29,1,0,0,0,189,199,5,15,0,0,190,192,5,30,0,0,191,
193,3,30,15,0,192,191,1,0,0,0,193,194,1,0,0,0,194,192,1,0,0,0,194,195,1,
0,0,0,195,196,1,0,0,0,196,197,5,31,0,0,197,199,1,0,0,0,198,189,1,0,0,0,198,
190,1,0,0,0,199,202,1,0,0,0,200,201,5,23,0,0,201,203,5,12,0,0,202,200,1,
0,0,0,202,203,1,0,0,0,203,31,1,0,0,0,204,205,5,5,0,0,205,206,3,8,4,0,206,
207,5,32,0,0,207,208,3,40,20,0,208,33,1,0,0,0,209,210,5,6,0,0,210,211,3,
8,4,0,211,215,5,35,0,0,212,214,3,36,18,0,213,212,1,0,0,0,214,217,1,0,0,0,
215,213,1,0,0,0,215,216,1,0,0,0,216,218,1,0,0,0,217,215,1,0,0,0,218,219,
5,10,0,0,219,35,1,0,0,0,220,221,3,8,4,0,221,222,3,40,20,0,222,223,5,35,0,
0,223,235,1,0,0,0,224,225,3,8,4,0,225,227,5,35,0,0,226,228,3,38,19,0,227,
226,1,0,0,0,228,229,1,0,0,0,229,227,1,0,0,0,229,230,1,0,0,0,230,231,1,0,
0,0,231,232,5,10,0,0,232,233,5,35,0,0,233,235,1,0,0,0,234,220,1,0,0,0,234,
224,1,0,0,0,235,37,1,0,0,0,236,238,3,8,4,0,237,239,3,40,20,0,238,237,1,0,
0,0,238,239,1,0,0,0,239,240,1,0,0,0,240,241,5,35,0,0,241,39,1,0,0,0,242,
244,3,42,21,0,243,242,1,0,0,0,244,245,1,0,0,0,245,243,1,0,0,0,245,246,1,
0,0,0,246,41,1,0,0,0,247,268,3,46,23,0,248,268,5,15,0,0,249,268,3,44,22,
0,250,255,3,8,4,0,251,252,5,33,0,0,252,254,3,8,4,0,253,251,1,0,0,0,254,257,
1,0,0,0,255,253,1,0,0,0,255,256,1,0,0,0,256,260,1,0,0,0,257,255,1,0,0,0,
258,259,5,34,0,0,259,261,5,12,0,0,260,258,1,0,0,0,260,261,1,0,0,0,261,268,
1,0,0,0,262,265,5,12,0,0,263,264,5,34,0,0,264,266,5,12,0,0,265,263,1,0,0,
0,265,266,1,0,0,0,266,268,1,0,0,0,267,247,1,0,0,0,267,248,1,0,0,0,267,249,
1,0,0,0,267,250,1,0,0,0,267,262,1,0,0,0,268,43,1,0,0,0,269,270,5,32,0,0,
270,273,5,12,0,0,271,272,5,34,0,0,272,274,5,12,0,0,273,271,1,0,0,0,273,274,
1,0,0,0,274,45,1,0,0,0,275,276,3,8,4,0,276,277,5,12,0,0,277,278,5,18,0,0,
278,282,5,12,0,0,279,280,3,8,4,0,280,281,5,12,0,0,281,283,1,0,0,0,282,279,
1,0,0,0,282,283,1,0,0,0,283,287,1,0,0,0,284,285,3,8,4,0,285,286,5,12,0,0,
286,288,1,0,0,0,287,284,1,0,0,0,287,288,1,0,0,0,288,290,1,0,0,0,289,291,
3,8,4,0,290,289,1,0,0,0,290,291,1,0,0,0,291,47,1,0,0,0,292,293,5,3,0,0,293,
295,5,12,0,0,294,296,3,8,4,0,295,294,1,0,0,0,295,296,1,0,0,0,296,297,1,0,
0,0,297,298,3,50,25,0,298,49,1,0,0,0,299,302,3,10,5,0,300,302,5,15,0,0,301,
299,1,0,0,0,301,300,1,0,0,0,302,303,1,0,0,0,303,316,3,52,26,0,304,305,3,
28,14,0,305,306,3,52,26,0,306,316,1,0,0,0,307,311,5,35,0,0,308,310,3,58,
29,0,309,308,1,0,0,0,310,313,1,0,0,0,311,309,1,0,0,0,311,312,1,0,0,0,312,
314,1,0,0,0,313,311,1,0,0,0,314,316,5,10,0,0,315,301,1,0,0,0,315,304,1,0,
0,0,315,307,1,0,0,0,316,51,1,0,0,0,317,319,5,35,0,0,318,317,1,0,0,0,318,
319,1,0,0,0,319,320,1,0,0,0,320,322,3,54,27,0,321,318,1,0,0,0,322,325,1,
0,0,0,323,321,1,0,0,0,323,324,1,0,0,0,324,332,1,0,0,0,325,323,1,0,0,0,326,
328,5,35,0,0,327,326,1,0,0,0,327,328,1,0,0,0,328,329,1,0,0,0,329,331,3,62,
31,0,330,327,1,0,0,0,331,334,1,0,0,0,332,330,1,0,0,0,332,333,1,0,0,0,333,
53,1,0,0,0,334,332,1,0,0,0,335,336,3,8,4,0,336,339,5,21,0,0,337,340,3,10,
5,0,338,340,3,56,28,0,339,337,1,0,0,0,339,338,1,0,0,0,340,55,1,0,0,0,341,
346,3,42,21,0,342,343,5,22,0,0,343,345,3,42,21,0,344,342,1,0,0,0,345,348,
1,0,0,0,346,344,1,0,0,0,346,347,1,0,0,0,347,57,1,0,0,0,348,346,1,0,0,0,349,
354,3,62,31,0,350,351,3,8,4,0,351,352,3,60,30,0,352,354,1,0,0,0,353,349,
1,0,0,0,353,350,1,0,0,0,354,355,1,0,0,0,355,356,5,35,0,0,356,59,1,0,0,0,
357,369,3,10,5,0,358,369,5,15,0,0,359,362,3,8,4,0,360,362,5,11,0,0,361,359,
1,0,0,0,361,360,1,0,0,0,362,363,1,0,0,0,363,369,3,10,5,0,364,366,5,12,0,
0,365,367,5,12,0,0,366,365,1,0,0,0,366,367,1,0,0,0,367,369,1,0,0,0,368,357,
1,0,0,0,368,358,1,0,0,0,368,361,1,0,0,0,368,364,1,0,0,0,369,61,1,0,0,0,370,
371,5,17,0,0,371,372,3,64,32,0,372,63,1,0,0,0,373,375,3,8,4,0,374,376,3,
66,33,0,375,374,1,0,0,0,375,376,1,0,0,0,376,378,1,0,0,0,377,379,3,68,34,
0,378,377,1,0,0,0,378,379,1,0,0,0,379,65,1,0,0,0,380,383,5,12,0,0,381,383,
3,10,5,0,382,380,1,0,0,0,382,381,1,0,0,0,383,67,1,0,0,0,384,390,3,64,32,
0,385,386,5,26,0,0,386,387,3,64,32,0,387,388,5,27,0,0,388,390,1,0,0,0,389,
384,1,0,0,0,389,385,1,0,0,0,390,69,1,0,0,0,391,392,5,7,0,0,392,396,5,35,
0,0,393,395,3,72,36,0,394,393,1,0,0,0,395,398,1,0,0,0,396,394,1,0,0,0,396,
397,1,0,0,0,397,399,1,0,0,0,398,396,1,0,0,0,399,400,5,10,0,0,400,71,1,0,
0,0,401,402,3,8,4,0,402,403,3,74,37,0,403,404,5,35,0,0,404,73,1,0,0,0,405,
406,3,8,4,0,406,407,3,8,4,0,407,408,5,35,0,0,408,413,3,76,38,0,409,410,5,
35,0,0,410,412,3,76,38,0,411,409,1,0,0,0,412,415,1,0,0,0,413,411,1,0,0,0,
413,414,1,0,0,0,414,426,1,0,0,0,415,413,1,0,0,0,416,417,3,8,4,0,417,421,
3,80,40,0,418,420,3,82,41,0,419,418,1,0,0,0,420,423,1,0,0,0,421,419,1,0,
0,0,421,422,1,0,0,0,422,426,1,0,0,0,423,421,1,0,0,0,424,426,3,84,42,0,425,
405,1,0,0,0,425,416,1,0,0,0,425,424,1,0,0,0,426,75,1,0,0,0,427,428,3,78,
39,0,428,438,5,21,0,0,429,430,3,8,4,0,430,434,3,80,40,0,431,433,3,82,41,
0,432,431,1,0,0,0,433,436,1,0,0,0,434,432,1,0,0,0,434,435,1,0,0,0,435,439,
1,0,0,0,436,434,1,0,0,0,437,439,3,84,42,0,438,429,1,0,0,0,438,437,1,0,0,
0,439,77,1,0,0,0,440,441,5,25,0,0,441,446,5,12,0,0,442,443,5,20,0,0,443,
446,5,12,0,0,444,446,3,8,4,0,445,440,1,0,0,0,445,442,1,0,0,0,445,444,1,0,
0,0,446,79,1,0,0,0,447,448,5,12,0,0,448,449,5,18,0,0,449,451,5,12,0,0,450,
447,1,0,0,0,450,451,1,0,0,0,451,81,1,0,0,0,452,453,7,4,0,0,453,457,5,12,
0,0,454,455,5,19,0,0,455,457,5,12,0,0,456,452,1,0,0,0,456,454,1,0,0,0,457,
83,1,0,0,0,458,465,5,12,0,0,459,461,5,22,0,0,460,459,1,0,0,0,460,461,1,0,
0,0,461,462,1,0,0,0,462,464,5,12,0,0,463,460,1,0,0,0,464,467,1,0,0,0,465,
463,1,0,0,0,465,466,1,0,0,0,466,85,1,0,0,0,467,465,1,0,0,0,58,91,105,111,
125,130,136,143,154,159,162,169,177,187,194,198,202,215,229,234,238,245,
255,260,265,267,273,282,287,290,295,301,311,315,318,323,327,332,339,346,
353,361,366,368,375,378,382,389,396,413,421,425,434,438,445,450,456,460,
465];


const atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

const decisionsToDFA = atn.decisionToState.map( (ds, index) => new antlr4.dfa.DFA(ds, index) );

const sharedContextCache = new antlr4.atn.PredictionContextCache();

export default class RasterklangDslParser extends antlr4.Parser {

    static grammarFileName = "RasterklangDsl.g4";
    static literalNames = [ null, "'tempo'", "'speed'", "'voice'", "'pattern'", 
                            "'table'", "'instrument'", "'filter'", "'target'", 
                            "'chip'", "'end'", null, null, null, null, null, 
                            null, "'|>'", "'..'", "'>>'", "'>='", "':'", 
                            "','", "'*'", "'/'", "'<'", "'('", "')'", "'{'", 
                            "'}'", "'['", "']'", "'='", "'+'", "'!'" ];
    static symbolicNames = [ null, "TEMPO", "SPEED", "VOICE", "PATTERN", 
                             "TABLE", "INSTRUMENT", "FILTER", "TARGET", 
                             "CHIP", "END", "HEX", "NUMBER", "DQSTRING", 
                             "BTSTRING", "SIGIL", "IDENT", "PIPE", "DOTDOT", 
                             "SHIFT", "GE", "COLON", "COMMA", "STAR", "SLASH", 
                             "LT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", 
                             "LBRACKET", "RBRACKET", "EQ", "PLUS", "BANG", 
                             "NEWLINE", "WS", "COMMENT" ];
    static ruleNames = [ "program", "statement", "unknownStmt", "unknownTail", 
                         "anyIdent", "str", "tempoStmt", "tempoValue", "tempoSetting", 
                         "targetStmt", "chipStmt", "targetSetting", "patternDecl", 
                         "patternEach", "arrangement", "arrangementElement", 
                         "tableDecl", "instrumentBlock", "instrumentField", 
                         "tableVerbLine", "tableRow", "tableToken", "absNote", 
                         "rampGenerator", "voiceStmt", "voiceBody", "voiceShortTail", 
                         "inlineAttr", "valueList", "voiceField", "voiceFieldValue", 
                         "pipelineStep", "transform", "transformArg", "transformRef", 
                         "filterBlock", "filterField", "filterFieldValue", 
                         "switchCase", "switchCond", "signalRangeOpt", "signalSpeedMod", 
                         "numberList" ];

    constructor(input) {
        super(input);
        this._interp = new antlr4.atn.ParserATNSimulator(this, atn, decisionsToDFA, sharedContextCache);
        this.ruleNames = RasterklangDslParser.ruleNames;
        this.literalNames = RasterklangDslParser.literalNames;
        this.symbolicNames = RasterklangDslParser.symbolicNames;
    }



	program() {
	    let localctx = new ProgramContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 0, RasterklangDslParser.RULE_program);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 91;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 66558) !== 0)) {
	            this.state = 86;
	            this.statement();
	            this.state = 87;
	            this.match(RasterklangDslParser.NEWLINE);
	            this.state = 93;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 94;
	        this.match(RasterklangDslParser.EOF);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	statement() {
	    let localctx = new StatementContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 2, RasterklangDslParser.RULE_statement);
	    try {
	        this.state = 105;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 1:
	        case 2:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 96;
	            this.tempoStmt();
	            break;
	        case 8:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 97;
	            this.targetStmt();
	            break;
	        case 9:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 98;
	            this.chipStmt();
	            break;
	        case 3:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 99;
	            this.voiceStmt();
	            break;
	        case 4:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 100;
	            this.patternDecl();
	            break;
	        case 5:
	            this.enterOuterAlt(localctx, 6);
	            this.state = 101;
	            this.tableDecl();
	            break;
	        case 6:
	            this.enterOuterAlt(localctx, 7);
	            this.state = 102;
	            this.instrumentBlock();
	            break;
	        case 7:
	            this.enterOuterAlt(localctx, 8);
	            this.state = 103;
	            this.filterBlock();
	            break;
	        case 16:
	            this.enterOuterAlt(localctx, 9);
	            this.state = 104;
	            this.unknownStmt();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	unknownStmt() {
	    let localctx = new UnknownStmtContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 4, RasterklangDslParser.RULE_unknownStmt);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 107;
	        this.match(RasterklangDslParser.IDENT);
	        this.state = 111;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 4294967294) !== 0) || ((((_la - 32)) & ~0x1f) === 0 && ((1 << (_la - 32)) & 55) !== 0)) {
	            this.state = 108;
	            this.unknownTail();
	            this.state = 113;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	unknownTail() {
	    let localctx = new UnknownTailContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 6, RasterklangDslParser.RULE_unknownTail);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 114;
	        _la = this._input.LA(1);
	        if(_la<=0 || _la===35) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	anyIdent() {
	    let localctx = new AnyIdentContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 8, RasterklangDslParser.RULE_anyIdent);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 116;
	        _la = this._input.LA(1);
	        if(!((((_la) & ~0x1f) === 0 && ((1 << _la) & 66558) !== 0))) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	str() {
	    let localctx = new StrContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 10, RasterklangDslParser.RULE_str);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 118;
	        _la = this._input.LA(1);
	        if(!(_la===13 || _la===14)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	tempoStmt() {
	    let localctx = new TempoStmtContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 12, RasterklangDslParser.RULE_tempoStmt);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 120;
	        _la = this._input.LA(1);
	        if(!(_la===1 || _la===2)) {
	        this._errHandler.recoverInline(this);
	        }
	        else {
	        	this._errHandler.reportMatch(this);
	            this.consume();
	        }
	        this.state = 121;
	        this.tempoValue();
	        this.state = 125;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 66558) !== 0)) {
	            this.state = 122;
	            this.tempoSetting();
	            this.state = 127;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	tempoValue() {
	    let localctx = new TempoValueContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 14, RasterklangDslParser.RULE_tempoValue);
	    try {
	        this.state = 130;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 12:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 128;
	            this.match(RasterklangDslParser.NUMBER);
	            break;
	        case 13:
	        case 14:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 129;
	            this.str();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	tempoSetting() {
	    let localctx = new TempoSettingContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 16, RasterklangDslParser.RULE_tempoSetting);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 132;
	        this.anyIdent();
	        this.state = 133;
	        this.match(RasterklangDslParser.NUMBER);
	        this.state = 136;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===24) {
	            this.state = 134;
	            this.match(RasterklangDslParser.SLASH);
	            this.state = 135;
	            this.match(RasterklangDslParser.NUMBER);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	targetStmt() {
	    let localctx = new TargetStmtContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 18, RasterklangDslParser.RULE_targetStmt);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 138;
	        this.match(RasterklangDslParser.TARGET);
	        this.state = 139;
	        this.match(RasterklangDslParser.LBRACE);
	        this.state = 143;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 66558) !== 0)) {
	            this.state = 140;
	            this.targetSetting();
	            this.state = 145;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 146;
	        this.match(RasterklangDslParser.RBRACE);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	chipStmt() {
	    let localctx = new ChipStmtContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 20, RasterklangDslParser.RULE_chipStmt);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 148;
	        this.match(RasterklangDslParser.CHIP);
	        this.state = 149;
	        this.match(RasterklangDslParser.NUMBER);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	targetSetting() {
	    let localctx = new TargetSettingContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 22, RasterklangDslParser.RULE_targetSetting);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 151;
	        this.anyIdent();
	        this.state = 154;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 12:
	            this.state = 152;
	            this.match(RasterklangDslParser.NUMBER);
	            break;
	        case 1:
	        case 2:
	        case 3:
	        case 4:
	        case 5:
	        case 6:
	        case 7:
	        case 8:
	        case 9:
	        case 16:
	            this.state = 153;
	            this.anyIdent();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	patternDecl() {
	    let localctx = new PatternDeclContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 24, RasterklangDslParser.RULE_patternDecl);
	    var _la = 0;
	    try {
	        this.state = 177;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,11,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 156;
	            this.match(RasterklangDslParser.PATTERN);
	            this.state = 157;
	            this.anyIdent();
	            this.state = 159;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===12) {
	                this.state = 158;
	                this.match(RasterklangDslParser.NUMBER);
	            }

	            this.state = 162;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if((((_la) & ~0x1f) === 0 && ((1 << _la) & 66558) !== 0)) {
	                this.state = 161;
	                this.patternEach();
	            }

	            this.state = 164;
	            this.match(RasterklangDslParser.EQ);
	            this.state = 165;
	            this.str();
	            this.state = 169;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while(_la===17) {
	                this.state = 166;
	                this.pipelineStep();
	                this.state = 171;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 172;
	            this.match(RasterklangDslParser.PATTERN);
	            this.state = 173;
	            this.anyIdent();
	            this.state = 174;
	            this.match(RasterklangDslParser.EQ);
	            this.state = 175;
	            this.arrangement();
	            break;

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	patternEach() {
	    let localctx = new PatternEachContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 26, RasterklangDslParser.RULE_patternEach);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 179;
	        this.anyIdent();
	        this.state = 180;
	        this.match(RasterklangDslParser.NUMBER);
	        this.state = 181;
	        this.match(RasterklangDslParser.SLASH);
	        this.state = 182;
	        this.match(RasterklangDslParser.NUMBER);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	arrangement() {
	    let localctx = new ArrangementContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 28, RasterklangDslParser.RULE_arrangement);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 185; 
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        do {
	            this.state = 184;
	            this.arrangementElement();
	            this.state = 187; 
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        } while(_la===15 || _la===30);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	arrangementElement() {
	    let localctx = new ArrangementElementContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 30, RasterklangDslParser.RULE_arrangementElement);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 198;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 15:
	            this.state = 189;
	            this.match(RasterklangDslParser.SIGIL);
	            break;
	        case 30:
	            this.state = 190;
	            this.match(RasterklangDslParser.LBRACKET);
	            this.state = 192; 
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            do {
	                this.state = 191;
	                this.arrangementElement();
	                this.state = 194; 
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            } while(_la===15 || _la===30);
	            this.state = 196;
	            this.match(RasterklangDslParser.RBRACKET);
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	        this.state = 202;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===23) {
	            this.state = 200;
	            this.match(RasterklangDslParser.STAR);
	            this.state = 201;
	            this.match(RasterklangDslParser.NUMBER);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	tableDecl() {
	    let localctx = new TableDeclContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 32, RasterklangDslParser.RULE_tableDecl);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 204;
	        this.match(RasterklangDslParser.TABLE);
	        this.state = 205;
	        this.anyIdent();
	        this.state = 206;
	        this.match(RasterklangDslParser.EQ);
	        this.state = 207;
	        this.tableRow();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	instrumentBlock() {
	    let localctx = new InstrumentBlockContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 34, RasterklangDslParser.RULE_instrumentBlock);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 209;
	        this.match(RasterklangDslParser.INSTRUMENT);
	        this.state = 210;
	        this.anyIdent();
	        this.state = 211;
	        this.match(RasterklangDslParser.NEWLINE);
	        this.state = 215;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 66558) !== 0)) {
	            this.state = 212;
	            this.instrumentField();
	            this.state = 217;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 218;
	        this.match(RasterklangDslParser.END);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	instrumentField() {
	    let localctx = new InstrumentFieldContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 36, RasterklangDslParser.RULE_instrumentField);
	    var _la = 0;
	    try {
	        this.state = 234;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,18,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 220;
	            this.anyIdent();
	            this.state = 221;
	            this.tableRow();
	            this.state = 222;
	            this.match(RasterklangDslParser.NEWLINE);
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 224;
	            this.anyIdent();
	            this.state = 225;
	            this.match(RasterklangDslParser.NEWLINE);
	            this.state = 227; 
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            do {
	                this.state = 226;
	                this.tableVerbLine();
	                this.state = 229; 
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            } while((((_la) & ~0x1f) === 0 && ((1 << _la) & 66558) !== 0));
	            this.state = 231;
	            this.match(RasterklangDslParser.END);
	            this.state = 232;
	            this.match(RasterklangDslParser.NEWLINE);
	            break;

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	tableVerbLine() {
	    let localctx = new TableVerbLineContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 38, RasterklangDslParser.RULE_tableVerbLine);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 236;
	        this.anyIdent();
	        this.state = 238;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(((((_la - 1)) & ~0x1f) === 0 && ((1 << (_la - 1)) & 2147535359) !== 0)) {
	            this.state = 237;
	            this.tableRow();
	        }

	        this.state = 240;
	        this.match(RasterklangDslParser.NEWLINE);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	tableRow() {
	    let localctx = new TableRowContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 40, RasterklangDslParser.RULE_tableRow);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 243; 
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        do {
	            this.state = 242;
	            this.tableToken();
	            this.state = 245; 
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        } while(((((_la - 1)) & ~0x1f) === 0 && ((1 << (_la - 1)) & 2147535359) !== 0));
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	tableToken() {
	    let localctx = new TableTokenContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 42, RasterklangDslParser.RULE_tableToken);
	    var _la = 0;
	    try {
	        this.state = 267;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,24,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 247;
	            this.rampGenerator();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 248;
	            this.match(RasterklangDslParser.SIGIL);
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 249;
	            this.absNote();
	            break;

	        case 4:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 250;
	            this.anyIdent();
	            this.state = 255;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while(_la===33) {
	                this.state = 251;
	                this.match(RasterklangDslParser.PLUS);
	                this.state = 252;
	                this.anyIdent();
	                this.state = 257;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            this.state = 260;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===34) {
	                this.state = 258;
	                this.match(RasterklangDslParser.BANG);
	                this.state = 259;
	                this.match(RasterklangDslParser.NUMBER);
	            }

	            break;

	        case 5:
	            this.enterOuterAlt(localctx, 5);
	            this.state = 262;
	            this.match(RasterklangDslParser.NUMBER);
	            this.state = 265;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===34) {
	                this.state = 263;
	                this.match(RasterklangDslParser.BANG);
	                this.state = 264;
	                this.match(RasterklangDslParser.NUMBER);
	            }

	            break;

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	absNote() {
	    let localctx = new AbsNoteContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 44, RasterklangDslParser.RULE_absNote);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 269;
	        this.match(RasterklangDslParser.EQ);
	        this.state = 270;
	        this.match(RasterklangDslParser.NUMBER);
	        this.state = 273;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===34) {
	            this.state = 271;
	            this.match(RasterklangDslParser.BANG);
	            this.state = 272;
	            this.match(RasterklangDslParser.NUMBER);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	rampGenerator() {
	    let localctx = new RampGeneratorContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 46, RasterklangDslParser.RULE_rampGenerator);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 275;
	        this.anyIdent();
	        this.state = 276;
	        this.match(RasterklangDslParser.NUMBER);
	        this.state = 277;
	        this.match(RasterklangDslParser.DOTDOT);
	        this.state = 278;
	        this.match(RasterklangDslParser.NUMBER);
	        this.state = 282;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,26,this._ctx);
	        if(la_===1) {
	            this.state = 279;
	            this.anyIdent();
	            this.state = 280;
	            this.match(RasterklangDslParser.NUMBER);

	        }
	        this.state = 287;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,27,this._ctx);
	        if(la_===1) {
	            this.state = 284;
	            this.anyIdent();
	            this.state = 285;
	            this.match(RasterklangDslParser.NUMBER);

	        }
	        this.state = 290;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,28,this._ctx);
	        if(la_===1) {
	            this.state = 289;
	            this.anyIdent();

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	voiceStmt() {
	    let localctx = new VoiceStmtContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 48, RasterklangDslParser.RULE_voiceStmt);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 292;
	        this.match(RasterklangDslParser.VOICE);
	        this.state = 293;
	        this.match(RasterklangDslParser.NUMBER);
	        this.state = 295;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if((((_la) & ~0x1f) === 0 && ((1 << _la) & 66558) !== 0)) {
	            this.state = 294;
	            this.anyIdent();
	        }

	        this.state = 297;
	        this.voiceBody();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	voiceBody() {
	    let localctx = new VoiceBodyContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 50, RasterklangDslParser.RULE_voiceBody);
	    var _la = 0;
	    try {
	        this.state = 315;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,32,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 301;
	            this._errHandler.sync(this);
	            switch(this._input.LA(1)) {
	            case 13:
	            case 14:
	                this.state = 299;
	                this.str();
	                break;
	            case 15:
	                this.state = 300;
	                this.match(RasterklangDslParser.SIGIL);
	                break;
	            default:
	                throw new antlr4.error.NoViableAltException(this);
	            }
	            this.state = 303;
	            this.voiceShortTail();
	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 304;
	            this.arrangement();
	            this.state = 305;
	            this.voiceShortTail();
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 307;
	            this.match(RasterklangDslParser.NEWLINE);
	            this.state = 311;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while((((_la) & ~0x1f) === 0 && ((1 << _la) & 197630) !== 0)) {
	                this.state = 308;
	                this.voiceField();
	                this.state = 313;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            this.state = 314;
	            this.match(RasterklangDslParser.END);
	            break;

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	voiceShortTail() {
	    let localctx = new VoiceShortTailContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 52, RasterklangDslParser.RULE_voiceShortTail);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 323;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,34,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 318;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                if(_la===35) {
	                    this.state = 317;
	                    this.match(RasterklangDslParser.NEWLINE);
	                }

	                this.state = 320;
	                this.inlineAttr(); 
	            }
	            this.state = 325;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,34,this._ctx);
	        }

	        this.state = 332;
	        this._errHandler.sync(this);
	        var _alt = this._interp.adaptivePredict(this._input,36,this._ctx)
	        while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	            if(_alt===1) {
	                this.state = 327;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	                if(_la===35) {
	                    this.state = 326;
	                    this.match(RasterklangDslParser.NEWLINE);
	                }

	                this.state = 329;
	                this.pipelineStep(); 
	            }
	            this.state = 334;
	            this._errHandler.sync(this);
	            _alt = this._interp.adaptivePredict(this._input,36,this._ctx);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	inlineAttr() {
	    let localctx = new InlineAttrContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 54, RasterklangDslParser.RULE_inlineAttr);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 335;
	        this.anyIdent();
	        this.state = 336;
	        this.match(RasterklangDslParser.COLON);
	        this.state = 339;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 13:
	        case 14:
	            this.state = 337;
	            this.str();
	            break;
	        case 1:
	        case 2:
	        case 3:
	        case 4:
	        case 5:
	        case 6:
	        case 7:
	        case 8:
	        case 9:
	        case 12:
	        case 15:
	        case 16:
	        case 32:
	            this.state = 338;
	            this.valueList();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	valueList() {
	    let localctx = new ValueListContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 56, RasterklangDslParser.RULE_valueList);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 341;
	        this.tableToken();
	        this.state = 346;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===22) {
	            this.state = 342;
	            this.match(RasterklangDslParser.COMMA);
	            this.state = 343;
	            this.tableToken();
	            this.state = 348;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	voiceField() {
	    let localctx = new VoiceFieldContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 58, RasterklangDslParser.RULE_voiceField);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 353;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 17:
	            this.state = 349;
	            this.pipelineStep();
	            break;
	        case 1:
	        case 2:
	        case 3:
	        case 4:
	        case 5:
	        case 6:
	        case 7:
	        case 8:
	        case 9:
	        case 16:
	            this.state = 350;
	            this.anyIdent();
	            this.state = 351;
	            this.voiceFieldValue();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	        this.state = 355;
	        this.match(RasterklangDslParser.NEWLINE);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	voiceFieldValue() {
	    let localctx = new VoiceFieldValueContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 60, RasterklangDslParser.RULE_voiceFieldValue);
	    var _la = 0;
	    try {
	        this.state = 368;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 13:
	        case 14:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 357;
	            this.str();
	            break;
	        case 15:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 358;
	            this.match(RasterklangDslParser.SIGIL);
	            break;
	        case 1:
	        case 2:
	        case 3:
	        case 4:
	        case 5:
	        case 6:
	        case 7:
	        case 8:
	        case 9:
	        case 11:
	        case 16:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 361;
	            this._errHandler.sync(this);
	            switch(this._input.LA(1)) {
	            case 1:
	            case 2:
	            case 3:
	            case 4:
	            case 5:
	            case 6:
	            case 7:
	            case 8:
	            case 9:
	            case 16:
	                this.state = 359;
	                this.anyIdent();
	                break;
	            case 11:
	                this.state = 360;
	                this.match(RasterklangDslParser.HEX);
	                break;
	            default:
	                throw new antlr4.error.NoViableAltException(this);
	            }
	            this.state = 363;
	            this.str();
	            break;
	        case 12:
	            this.enterOuterAlt(localctx, 4);
	            this.state = 364;
	            this.match(RasterklangDslParser.NUMBER);
	            this.state = 366;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===12) {
	                this.state = 365;
	                this.match(RasterklangDslParser.NUMBER);
	            }

	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	pipelineStep() {
	    let localctx = new PipelineStepContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 62, RasterklangDslParser.RULE_pipelineStep);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 370;
	        this.match(RasterklangDslParser.PIPE);
	        this.state = 371;
	        this.transform();
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	transform() {
	    let localctx = new TransformContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 64, RasterklangDslParser.RULE_transform);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 373;
	        this.anyIdent();
	        this.state = 375;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if((((_la) & ~0x1f) === 0 && ((1 << _la) & 28672) !== 0)) {
	            this.state = 374;
	            this.transformArg();
	        }

	        this.state = 378;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if((((_la) & ~0x1f) === 0 && ((1 << _la) & 67175422) !== 0)) {
	            this.state = 377;
	            this.transformRef();
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	transformArg() {
	    let localctx = new TransformArgContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 66, RasterklangDslParser.RULE_transformArg);
	    try {
	        this.state = 382;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 12:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 380;
	            this.match(RasterklangDslParser.NUMBER);
	            break;
	        case 13:
	        case 14:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 381;
	            this.str();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	transformRef() {
	    let localctx = new TransformRefContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 68, RasterklangDslParser.RULE_transformRef);
	    try {
	        this.state = 389;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 1:
	        case 2:
	        case 3:
	        case 4:
	        case 5:
	        case 6:
	        case 7:
	        case 8:
	        case 9:
	        case 16:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 384;
	            this.transform();
	            break;
	        case 26:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 385;
	            this.match(RasterklangDslParser.LPAREN);
	            this.state = 386;
	            this.transform();
	            this.state = 387;
	            this.match(RasterklangDslParser.RPAREN);
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	filterBlock() {
	    let localctx = new FilterBlockContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 70, RasterklangDslParser.RULE_filterBlock);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 391;
	        this.match(RasterklangDslParser.FILTER);
	        this.state = 392;
	        this.match(RasterklangDslParser.NEWLINE);
	        this.state = 396;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while((((_la) & ~0x1f) === 0 && ((1 << _la) & 66558) !== 0)) {
	            this.state = 393;
	            this.filterField();
	            this.state = 398;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	        this.state = 399;
	        this.match(RasterklangDslParser.END);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	filterField() {
	    let localctx = new FilterFieldContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 72, RasterklangDslParser.RULE_filterField);
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 401;
	        this.anyIdent();
	        this.state = 402;
	        this.filterFieldValue();
	        this.state = 403;
	        this.match(RasterklangDslParser.NEWLINE);
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	filterFieldValue() {
	    let localctx = new FilterFieldValueContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 74, RasterklangDslParser.RULE_filterFieldValue);
	    var _la = 0;
	    try {
	        this.state = 425;
	        this._errHandler.sync(this);
	        var la_ = this._interp.adaptivePredict(this._input,50,this._ctx);
	        switch(la_) {
	        case 1:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 405;
	            this.anyIdent();
	            this.state = 406;
	            this.anyIdent();
	            this.state = 407;
	            this.match(RasterklangDslParser.NEWLINE);
	            this.state = 408;
	            this.switchCase();
	            this.state = 413;
	            this._errHandler.sync(this);
	            var _alt = this._interp.adaptivePredict(this._input,48,this._ctx)
	            while(_alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER) {
	                if(_alt===1) {
	                    this.state = 409;
	                    this.match(RasterklangDslParser.NEWLINE);
	                    this.state = 410;
	                    this.switchCase(); 
	                }
	                this.state = 415;
	                this._errHandler.sync(this);
	                _alt = this._interp.adaptivePredict(this._input,48,this._ctx);
	            }

	            break;

	        case 2:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 416;
	            this.anyIdent();
	            this.state = 417;
	            this.signalRangeOpt();
	            this.state = 421;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while((((_la) & ~0x1f) === 0 && ((1 << _la) & 25690112) !== 0)) {
	                this.state = 418;
	                this.signalSpeedMod();
	                this.state = 423;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            break;

	        case 3:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 424;
	            this.numberList();
	            break;

	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	switchCase() {
	    let localctx = new SwitchCaseContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 76, RasterklangDslParser.RULE_switchCase);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 427;
	        this.switchCond();
	        this.state = 428;
	        this.match(RasterklangDslParser.COLON);
	        this.state = 438;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 1:
	        case 2:
	        case 3:
	        case 4:
	        case 5:
	        case 6:
	        case 7:
	        case 8:
	        case 9:
	        case 16:
	            this.state = 429;
	            this.anyIdent();
	            this.state = 430;
	            this.signalRangeOpt();
	            this.state = 434;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            while((((_la) & ~0x1f) === 0 && ((1 << _la) & 25690112) !== 0)) {
	                this.state = 431;
	                this.signalSpeedMod();
	                this.state = 436;
	                this._errHandler.sync(this);
	                _la = this._input.LA(1);
	            }
	            break;
	        case 12:
	            this.state = 437;
	            this.numberList();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	switchCond() {
	    let localctx = new SwitchCondContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 78, RasterklangDslParser.RULE_switchCond);
	    try {
	        this.state = 445;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 25:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 440;
	            this.match(RasterklangDslParser.LT);
	            this.state = 441;
	            this.match(RasterklangDslParser.NUMBER);
	            break;
	        case 20:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 442;
	            this.match(RasterklangDslParser.GE);
	            this.state = 443;
	            this.match(RasterklangDslParser.NUMBER);
	            break;
	        case 1:
	        case 2:
	        case 3:
	        case 4:
	        case 5:
	        case 6:
	        case 7:
	        case 8:
	        case 9:
	        case 16:
	            this.enterOuterAlt(localctx, 3);
	            this.state = 444;
	            this.anyIdent();
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	signalRangeOpt() {
	    let localctx = new SignalRangeOptContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 80, RasterklangDslParser.RULE_signalRangeOpt);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 450;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        if(_la===12) {
	            this.state = 447;
	            this.match(RasterklangDslParser.NUMBER);
	            this.state = 448;
	            this.match(RasterklangDslParser.DOTDOT);
	            this.state = 449;
	            this.match(RasterklangDslParser.NUMBER);
	        }

	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	signalSpeedMod() {
	    let localctx = new SignalSpeedModContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 82, RasterklangDslParser.RULE_signalSpeedMod);
	    var _la = 0;
	    try {
	        this.state = 456;
	        this._errHandler.sync(this);
	        switch(this._input.LA(1)) {
	        case 23:
	        case 24:
	            this.enterOuterAlt(localctx, 1);
	            this.state = 452;
	            _la = this._input.LA(1);
	            if(!(_la===23 || _la===24)) {
	            this._errHandler.recoverInline(this);
	            }
	            else {
	            	this._errHandler.reportMatch(this);
	                this.consume();
	            }
	            this.state = 453;
	            this.match(RasterklangDslParser.NUMBER);
	            break;
	        case 19:
	            this.enterOuterAlt(localctx, 2);
	            this.state = 454;
	            this.match(RasterklangDslParser.SHIFT);
	            this.state = 455;
	            this.match(RasterklangDslParser.NUMBER);
	            break;
	        default:
	            throw new antlr4.error.NoViableAltException(this);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}



	numberList() {
	    let localctx = new NumberListContext(this, this._ctx, this.state);
	    this.enterRule(localctx, 84, RasterklangDslParser.RULE_numberList);
	    var _la = 0;
	    try {
	        this.enterOuterAlt(localctx, 1);
	        this.state = 458;
	        this.match(RasterklangDslParser.NUMBER);
	        this.state = 465;
	        this._errHandler.sync(this);
	        _la = this._input.LA(1);
	        while(_la===12 || _la===22) {
	            this.state = 460;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	            if(_la===22) {
	                this.state = 459;
	                this.match(RasterklangDslParser.COMMA);
	            }

	            this.state = 462;
	            this.match(RasterklangDslParser.NUMBER);
	            this.state = 467;
	            this._errHandler.sync(this);
	            _la = this._input.LA(1);
	        }
	    } catch (re) {
	    	if(re instanceof antlr4.error.RecognitionException) {
		        localctx.exception = re;
		        this._errHandler.reportError(this, re);
		        this._errHandler.recover(this, re);
		    } else {
		    	throw re;
		    }
	    } finally {
	        this.exitRule();
	    }
	    return localctx;
	}


}

RasterklangDslParser.EOF = antlr4.Token.EOF;
RasterklangDslParser.TEMPO = 1;
RasterklangDslParser.SPEED = 2;
RasterklangDslParser.VOICE = 3;
RasterklangDslParser.PATTERN = 4;
RasterklangDslParser.TABLE = 5;
RasterklangDslParser.INSTRUMENT = 6;
RasterklangDslParser.FILTER = 7;
RasterklangDslParser.TARGET = 8;
RasterklangDslParser.CHIP = 9;
RasterklangDslParser.END = 10;
RasterklangDslParser.HEX = 11;
RasterklangDslParser.NUMBER = 12;
RasterklangDslParser.DQSTRING = 13;
RasterklangDslParser.BTSTRING = 14;
RasterklangDslParser.SIGIL = 15;
RasterklangDslParser.IDENT = 16;
RasterklangDslParser.PIPE = 17;
RasterklangDslParser.DOTDOT = 18;
RasterklangDslParser.SHIFT = 19;
RasterklangDslParser.GE = 20;
RasterklangDslParser.COLON = 21;
RasterklangDslParser.COMMA = 22;
RasterklangDslParser.STAR = 23;
RasterklangDslParser.SLASH = 24;
RasterklangDslParser.LT = 25;
RasterklangDslParser.LPAREN = 26;
RasterklangDslParser.RPAREN = 27;
RasterklangDslParser.LBRACE = 28;
RasterklangDslParser.RBRACE = 29;
RasterklangDslParser.LBRACKET = 30;
RasterklangDslParser.RBRACKET = 31;
RasterklangDslParser.EQ = 32;
RasterklangDslParser.PLUS = 33;
RasterklangDslParser.BANG = 34;
RasterklangDslParser.NEWLINE = 35;
RasterklangDslParser.WS = 36;
RasterklangDslParser.COMMENT = 37;

RasterklangDslParser.RULE_program = 0;
RasterklangDslParser.RULE_statement = 1;
RasterklangDslParser.RULE_unknownStmt = 2;
RasterklangDslParser.RULE_unknownTail = 3;
RasterklangDslParser.RULE_anyIdent = 4;
RasterklangDslParser.RULE_str = 5;
RasterklangDslParser.RULE_tempoStmt = 6;
RasterklangDslParser.RULE_tempoValue = 7;
RasterklangDslParser.RULE_tempoSetting = 8;
RasterklangDslParser.RULE_targetStmt = 9;
RasterklangDslParser.RULE_chipStmt = 10;
RasterklangDslParser.RULE_targetSetting = 11;
RasterklangDslParser.RULE_patternDecl = 12;
RasterklangDslParser.RULE_patternEach = 13;
RasterklangDslParser.RULE_arrangement = 14;
RasterklangDslParser.RULE_arrangementElement = 15;
RasterklangDslParser.RULE_tableDecl = 16;
RasterklangDslParser.RULE_instrumentBlock = 17;
RasterklangDslParser.RULE_instrumentField = 18;
RasterklangDslParser.RULE_tableVerbLine = 19;
RasterklangDslParser.RULE_tableRow = 20;
RasterklangDslParser.RULE_tableToken = 21;
RasterklangDslParser.RULE_absNote = 22;
RasterklangDslParser.RULE_rampGenerator = 23;
RasterklangDslParser.RULE_voiceStmt = 24;
RasterklangDslParser.RULE_voiceBody = 25;
RasterklangDslParser.RULE_voiceShortTail = 26;
RasterklangDslParser.RULE_inlineAttr = 27;
RasterklangDslParser.RULE_valueList = 28;
RasterklangDslParser.RULE_voiceField = 29;
RasterklangDslParser.RULE_voiceFieldValue = 30;
RasterklangDslParser.RULE_pipelineStep = 31;
RasterklangDslParser.RULE_transform = 32;
RasterklangDslParser.RULE_transformArg = 33;
RasterklangDslParser.RULE_transformRef = 34;
RasterklangDslParser.RULE_filterBlock = 35;
RasterklangDslParser.RULE_filterField = 36;
RasterklangDslParser.RULE_filterFieldValue = 37;
RasterklangDslParser.RULE_switchCase = 38;
RasterklangDslParser.RULE_switchCond = 39;
RasterklangDslParser.RULE_signalRangeOpt = 40;
RasterklangDslParser.RULE_signalSpeedMod = 41;
RasterklangDslParser.RULE_numberList = 42;

class ProgramContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_program;
    }

	EOF() {
	    return this.getToken(RasterklangDslParser.EOF, 0);
	};

	statement = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(StatementContext);
	    } else {
	        return this.getTypedRuleContext(StatementContext,i);
	    }
	};

	NEWLINE = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NEWLINE);
	    } else {
	        return this.getToken(RasterklangDslParser.NEWLINE, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterProgram(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitProgram(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitProgram(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class StatementContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_statement;
    }

	tempoStmt() {
	    return this.getTypedRuleContext(TempoStmtContext,0);
	};

	targetStmt() {
	    return this.getTypedRuleContext(TargetStmtContext,0);
	};

	chipStmt() {
	    return this.getTypedRuleContext(ChipStmtContext,0);
	};

	voiceStmt() {
	    return this.getTypedRuleContext(VoiceStmtContext,0);
	};

	patternDecl() {
	    return this.getTypedRuleContext(PatternDeclContext,0);
	};

	tableDecl() {
	    return this.getTypedRuleContext(TableDeclContext,0);
	};

	instrumentBlock() {
	    return this.getTypedRuleContext(InstrumentBlockContext,0);
	};

	filterBlock() {
	    return this.getTypedRuleContext(FilterBlockContext,0);
	};

	unknownStmt() {
	    return this.getTypedRuleContext(UnknownStmtContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterStatement(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitStatement(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitStatement(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class UnknownStmtContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_unknownStmt;
    }

	IDENT() {
	    return this.getToken(RasterklangDslParser.IDENT, 0);
	};

	unknownTail = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(UnknownTailContext);
	    } else {
	        return this.getTypedRuleContext(UnknownTailContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterUnknownStmt(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitUnknownStmt(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitUnknownStmt(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class UnknownTailContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_unknownTail;
    }

	NEWLINE() {
	    return this.getToken(RasterklangDslParser.NEWLINE, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterUnknownTail(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitUnknownTail(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitUnknownTail(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class AnyIdentContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_anyIdent;
    }

	IDENT() {
	    return this.getToken(RasterklangDslParser.IDENT, 0);
	};

	TEMPO() {
	    return this.getToken(RasterklangDslParser.TEMPO, 0);
	};

	SPEED() {
	    return this.getToken(RasterklangDslParser.SPEED, 0);
	};

	VOICE() {
	    return this.getToken(RasterklangDslParser.VOICE, 0);
	};

	PATTERN() {
	    return this.getToken(RasterklangDslParser.PATTERN, 0);
	};

	TABLE() {
	    return this.getToken(RasterklangDslParser.TABLE, 0);
	};

	INSTRUMENT() {
	    return this.getToken(RasterklangDslParser.INSTRUMENT, 0);
	};

	FILTER() {
	    return this.getToken(RasterklangDslParser.FILTER, 0);
	};

	TARGET() {
	    return this.getToken(RasterklangDslParser.TARGET, 0);
	};

	CHIP() {
	    return this.getToken(RasterklangDslParser.CHIP, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterAnyIdent(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitAnyIdent(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitAnyIdent(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class StrContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_str;
    }

	DQSTRING() {
	    return this.getToken(RasterklangDslParser.DQSTRING, 0);
	};

	BTSTRING() {
	    return this.getToken(RasterklangDslParser.BTSTRING, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterStr(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitStr(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitStr(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TempoStmtContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_tempoStmt;
    }

	tempoValue() {
	    return this.getTypedRuleContext(TempoValueContext,0);
	};

	TEMPO() {
	    return this.getToken(RasterklangDslParser.TEMPO, 0);
	};

	SPEED() {
	    return this.getToken(RasterklangDslParser.SPEED, 0);
	};

	tempoSetting = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(TempoSettingContext);
	    } else {
	        return this.getTypedRuleContext(TempoSettingContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTempoStmt(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTempoStmt(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTempoStmt(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TempoValueContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_tempoValue;
    }

	NUMBER() {
	    return this.getToken(RasterklangDslParser.NUMBER, 0);
	};

	str() {
	    return this.getTypedRuleContext(StrContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTempoValue(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTempoValue(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTempoValue(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TempoSettingContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_tempoSetting;
    }

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	NUMBER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NUMBER);
	    } else {
	        return this.getToken(RasterklangDslParser.NUMBER, i);
	    }
	};


	SLASH() {
	    return this.getToken(RasterklangDslParser.SLASH, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTempoSetting(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTempoSetting(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTempoSetting(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TargetStmtContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_targetStmt;
    }

	TARGET() {
	    return this.getToken(RasterklangDslParser.TARGET, 0);
	};

	LBRACE() {
	    return this.getToken(RasterklangDslParser.LBRACE, 0);
	};

	RBRACE() {
	    return this.getToken(RasterklangDslParser.RBRACE, 0);
	};

	targetSetting = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(TargetSettingContext);
	    } else {
	        return this.getTypedRuleContext(TargetSettingContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTargetStmt(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTargetStmt(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTargetStmt(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class ChipStmtContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_chipStmt;
    }

	CHIP() {
	    return this.getToken(RasterklangDslParser.CHIP, 0);
	};

	NUMBER() {
	    return this.getToken(RasterklangDslParser.NUMBER, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterChipStmt(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitChipStmt(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitChipStmt(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TargetSettingContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_targetSetting;
    }

	anyIdent = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(AnyIdentContext);
	    } else {
	        return this.getTypedRuleContext(AnyIdentContext,i);
	    }
	};

	NUMBER() {
	    return this.getToken(RasterklangDslParser.NUMBER, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTargetSetting(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTargetSetting(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTargetSetting(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class PatternDeclContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_patternDecl;
    }

	PATTERN() {
	    return this.getToken(RasterklangDslParser.PATTERN, 0);
	};

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	EQ() {
	    return this.getToken(RasterklangDslParser.EQ, 0);
	};

	str() {
	    return this.getTypedRuleContext(StrContext,0);
	};

	NUMBER() {
	    return this.getToken(RasterklangDslParser.NUMBER, 0);
	};

	patternEach() {
	    return this.getTypedRuleContext(PatternEachContext,0);
	};

	pipelineStep = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(PipelineStepContext);
	    } else {
	        return this.getTypedRuleContext(PipelineStepContext,i);
	    }
	};

	arrangement() {
	    return this.getTypedRuleContext(ArrangementContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterPatternDecl(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitPatternDecl(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitPatternDecl(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class PatternEachContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_patternEach;
    }

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	NUMBER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NUMBER);
	    } else {
	        return this.getToken(RasterklangDslParser.NUMBER, i);
	    }
	};


	SLASH() {
	    return this.getToken(RasterklangDslParser.SLASH, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterPatternEach(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitPatternEach(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitPatternEach(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class ArrangementContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_arrangement;
    }

	arrangementElement = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(ArrangementElementContext);
	    } else {
	        return this.getTypedRuleContext(ArrangementElementContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterArrangement(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitArrangement(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitArrangement(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class ArrangementElementContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_arrangementElement;
    }

	SIGIL() {
	    return this.getToken(RasterklangDslParser.SIGIL, 0);
	};

	LBRACKET() {
	    return this.getToken(RasterklangDslParser.LBRACKET, 0);
	};

	RBRACKET() {
	    return this.getToken(RasterklangDslParser.RBRACKET, 0);
	};

	STAR() {
	    return this.getToken(RasterklangDslParser.STAR, 0);
	};

	NUMBER() {
	    return this.getToken(RasterklangDslParser.NUMBER, 0);
	};

	arrangementElement = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(ArrangementElementContext);
	    } else {
	        return this.getTypedRuleContext(ArrangementElementContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterArrangementElement(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitArrangementElement(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitArrangementElement(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TableDeclContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_tableDecl;
    }

	TABLE() {
	    return this.getToken(RasterklangDslParser.TABLE, 0);
	};

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	EQ() {
	    return this.getToken(RasterklangDslParser.EQ, 0);
	};

	tableRow() {
	    return this.getTypedRuleContext(TableRowContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTableDecl(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTableDecl(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTableDecl(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class InstrumentBlockContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_instrumentBlock;
    }

	INSTRUMENT() {
	    return this.getToken(RasterklangDslParser.INSTRUMENT, 0);
	};

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	NEWLINE() {
	    return this.getToken(RasterklangDslParser.NEWLINE, 0);
	};

	END() {
	    return this.getToken(RasterklangDslParser.END, 0);
	};

	instrumentField = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(InstrumentFieldContext);
	    } else {
	        return this.getTypedRuleContext(InstrumentFieldContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterInstrumentBlock(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitInstrumentBlock(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitInstrumentBlock(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class InstrumentFieldContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_instrumentField;
    }

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	tableRow() {
	    return this.getTypedRuleContext(TableRowContext,0);
	};

	NEWLINE = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NEWLINE);
	    } else {
	        return this.getToken(RasterklangDslParser.NEWLINE, i);
	    }
	};


	END() {
	    return this.getToken(RasterklangDslParser.END, 0);
	};

	tableVerbLine = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(TableVerbLineContext);
	    } else {
	        return this.getTypedRuleContext(TableVerbLineContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterInstrumentField(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitInstrumentField(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitInstrumentField(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TableVerbLineContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_tableVerbLine;
    }

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	NEWLINE() {
	    return this.getToken(RasterklangDslParser.NEWLINE, 0);
	};

	tableRow() {
	    return this.getTypedRuleContext(TableRowContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTableVerbLine(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTableVerbLine(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTableVerbLine(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TableRowContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_tableRow;
    }

	tableToken = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(TableTokenContext);
	    } else {
	        return this.getTypedRuleContext(TableTokenContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTableRow(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTableRow(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTableRow(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TableTokenContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_tableToken;
    }

	rampGenerator() {
	    return this.getTypedRuleContext(RampGeneratorContext,0);
	};

	SIGIL() {
	    return this.getToken(RasterklangDslParser.SIGIL, 0);
	};

	absNote() {
	    return this.getTypedRuleContext(AbsNoteContext,0);
	};

	anyIdent = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(AnyIdentContext);
	    } else {
	        return this.getTypedRuleContext(AnyIdentContext,i);
	    }
	};

	PLUS = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.PLUS);
	    } else {
	        return this.getToken(RasterklangDslParser.PLUS, i);
	    }
	};


	BANG() {
	    return this.getToken(RasterklangDslParser.BANG, 0);
	};

	NUMBER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NUMBER);
	    } else {
	        return this.getToken(RasterklangDslParser.NUMBER, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTableToken(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTableToken(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTableToken(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class AbsNoteContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_absNote;
    }

	EQ() {
	    return this.getToken(RasterklangDslParser.EQ, 0);
	};

	NUMBER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NUMBER);
	    } else {
	        return this.getToken(RasterklangDslParser.NUMBER, i);
	    }
	};


	BANG() {
	    return this.getToken(RasterklangDslParser.BANG, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterAbsNote(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitAbsNote(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitAbsNote(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class RampGeneratorContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_rampGenerator;
    }

	anyIdent = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(AnyIdentContext);
	    } else {
	        return this.getTypedRuleContext(AnyIdentContext,i);
	    }
	};

	NUMBER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NUMBER);
	    } else {
	        return this.getToken(RasterklangDslParser.NUMBER, i);
	    }
	};


	DOTDOT() {
	    return this.getToken(RasterklangDslParser.DOTDOT, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterRampGenerator(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitRampGenerator(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitRampGenerator(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class VoiceStmtContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_voiceStmt;
    }

	VOICE() {
	    return this.getToken(RasterklangDslParser.VOICE, 0);
	};

	NUMBER() {
	    return this.getToken(RasterklangDslParser.NUMBER, 0);
	};

	voiceBody() {
	    return this.getTypedRuleContext(VoiceBodyContext,0);
	};

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterVoiceStmt(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitVoiceStmt(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitVoiceStmt(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class VoiceBodyContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_voiceBody;
    }

	voiceShortTail() {
	    return this.getTypedRuleContext(VoiceShortTailContext,0);
	};

	str() {
	    return this.getTypedRuleContext(StrContext,0);
	};

	SIGIL() {
	    return this.getToken(RasterklangDslParser.SIGIL, 0);
	};

	arrangement() {
	    return this.getTypedRuleContext(ArrangementContext,0);
	};

	NEWLINE() {
	    return this.getToken(RasterklangDslParser.NEWLINE, 0);
	};

	END() {
	    return this.getToken(RasterklangDslParser.END, 0);
	};

	voiceField = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(VoiceFieldContext);
	    } else {
	        return this.getTypedRuleContext(VoiceFieldContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterVoiceBody(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitVoiceBody(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitVoiceBody(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class VoiceShortTailContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_voiceShortTail;
    }

	inlineAttr = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(InlineAttrContext);
	    } else {
	        return this.getTypedRuleContext(InlineAttrContext,i);
	    }
	};

	pipelineStep = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(PipelineStepContext);
	    } else {
	        return this.getTypedRuleContext(PipelineStepContext,i);
	    }
	};

	NEWLINE = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NEWLINE);
	    } else {
	        return this.getToken(RasterklangDslParser.NEWLINE, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterVoiceShortTail(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitVoiceShortTail(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitVoiceShortTail(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class InlineAttrContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_inlineAttr;
    }

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	COLON() {
	    return this.getToken(RasterklangDslParser.COLON, 0);
	};

	str() {
	    return this.getTypedRuleContext(StrContext,0);
	};

	valueList() {
	    return this.getTypedRuleContext(ValueListContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterInlineAttr(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitInlineAttr(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitInlineAttr(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class ValueListContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_valueList;
    }

	tableToken = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(TableTokenContext);
	    } else {
	        return this.getTypedRuleContext(TableTokenContext,i);
	    }
	};

	COMMA = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.COMMA);
	    } else {
	        return this.getToken(RasterklangDslParser.COMMA, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterValueList(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitValueList(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitValueList(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class VoiceFieldContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_voiceField;
    }

	NEWLINE() {
	    return this.getToken(RasterklangDslParser.NEWLINE, 0);
	};

	pipelineStep() {
	    return this.getTypedRuleContext(PipelineStepContext,0);
	};

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	voiceFieldValue() {
	    return this.getTypedRuleContext(VoiceFieldValueContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterVoiceField(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitVoiceField(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitVoiceField(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class VoiceFieldValueContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_voiceFieldValue;
    }

	str() {
	    return this.getTypedRuleContext(StrContext,0);
	};

	SIGIL() {
	    return this.getToken(RasterklangDslParser.SIGIL, 0);
	};

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	HEX() {
	    return this.getToken(RasterklangDslParser.HEX, 0);
	};

	NUMBER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NUMBER);
	    } else {
	        return this.getToken(RasterklangDslParser.NUMBER, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterVoiceFieldValue(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitVoiceFieldValue(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitVoiceFieldValue(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class PipelineStepContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_pipelineStep;
    }

	PIPE() {
	    return this.getToken(RasterklangDslParser.PIPE, 0);
	};

	transform() {
	    return this.getTypedRuleContext(TransformContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterPipelineStep(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitPipelineStep(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitPipelineStep(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TransformContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_transform;
    }

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	transformArg() {
	    return this.getTypedRuleContext(TransformArgContext,0);
	};

	transformRef() {
	    return this.getTypedRuleContext(TransformRefContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTransform(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTransform(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTransform(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TransformArgContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_transformArg;
    }

	NUMBER() {
	    return this.getToken(RasterklangDslParser.NUMBER, 0);
	};

	str() {
	    return this.getTypedRuleContext(StrContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTransformArg(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTransformArg(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTransformArg(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class TransformRefContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_transformRef;
    }

	transform() {
	    return this.getTypedRuleContext(TransformContext,0);
	};

	LPAREN() {
	    return this.getToken(RasterklangDslParser.LPAREN, 0);
	};

	RPAREN() {
	    return this.getToken(RasterklangDslParser.RPAREN, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterTransformRef(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitTransformRef(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitTransformRef(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class FilterBlockContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_filterBlock;
    }

	FILTER() {
	    return this.getToken(RasterklangDslParser.FILTER, 0);
	};

	NEWLINE() {
	    return this.getToken(RasterklangDslParser.NEWLINE, 0);
	};

	END() {
	    return this.getToken(RasterklangDslParser.END, 0);
	};

	filterField = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(FilterFieldContext);
	    } else {
	        return this.getTypedRuleContext(FilterFieldContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterFilterBlock(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitFilterBlock(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitFilterBlock(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class FilterFieldContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_filterField;
    }

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	filterFieldValue() {
	    return this.getTypedRuleContext(FilterFieldValueContext,0);
	};

	NEWLINE() {
	    return this.getToken(RasterklangDslParser.NEWLINE, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterFilterField(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitFilterField(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitFilterField(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class FilterFieldValueContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_filterFieldValue;
    }

	anyIdent = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(AnyIdentContext);
	    } else {
	        return this.getTypedRuleContext(AnyIdentContext,i);
	    }
	};

	NEWLINE = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NEWLINE);
	    } else {
	        return this.getToken(RasterklangDslParser.NEWLINE, i);
	    }
	};


	switchCase = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(SwitchCaseContext);
	    } else {
	        return this.getTypedRuleContext(SwitchCaseContext,i);
	    }
	};

	signalRangeOpt() {
	    return this.getTypedRuleContext(SignalRangeOptContext,0);
	};

	signalSpeedMod = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(SignalSpeedModContext);
	    } else {
	        return this.getTypedRuleContext(SignalSpeedModContext,i);
	    }
	};

	numberList() {
	    return this.getTypedRuleContext(NumberListContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterFilterFieldValue(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitFilterFieldValue(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitFilterFieldValue(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class SwitchCaseContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_switchCase;
    }

	switchCond() {
	    return this.getTypedRuleContext(SwitchCondContext,0);
	};

	COLON() {
	    return this.getToken(RasterklangDslParser.COLON, 0);
	};

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	signalRangeOpt() {
	    return this.getTypedRuleContext(SignalRangeOptContext,0);
	};

	numberList() {
	    return this.getTypedRuleContext(NumberListContext,0);
	};

	signalSpeedMod = function(i) {
	    if(i===undefined) {
	        i = null;
	    }
	    if(i===null) {
	        return this.getTypedRuleContexts(SignalSpeedModContext);
	    } else {
	        return this.getTypedRuleContext(SignalSpeedModContext,i);
	    }
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterSwitchCase(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitSwitchCase(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitSwitchCase(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class SwitchCondContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_switchCond;
    }

	LT() {
	    return this.getToken(RasterklangDslParser.LT, 0);
	};

	NUMBER() {
	    return this.getToken(RasterklangDslParser.NUMBER, 0);
	};

	GE() {
	    return this.getToken(RasterklangDslParser.GE, 0);
	};

	anyIdent() {
	    return this.getTypedRuleContext(AnyIdentContext,0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterSwitchCond(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitSwitchCond(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitSwitchCond(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class SignalRangeOptContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_signalRangeOpt;
    }

	NUMBER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NUMBER);
	    } else {
	        return this.getToken(RasterklangDslParser.NUMBER, i);
	    }
	};


	DOTDOT() {
	    return this.getToken(RasterklangDslParser.DOTDOT, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterSignalRangeOpt(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitSignalRangeOpt(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitSignalRangeOpt(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class SignalSpeedModContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_signalSpeedMod;
    }

	NUMBER() {
	    return this.getToken(RasterklangDslParser.NUMBER, 0);
	};

	STAR() {
	    return this.getToken(RasterklangDslParser.STAR, 0);
	};

	SLASH() {
	    return this.getToken(RasterklangDslParser.SLASH, 0);
	};

	SHIFT() {
	    return this.getToken(RasterklangDslParser.SHIFT, 0);
	};

	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterSignalSpeedMod(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitSignalSpeedMod(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitSignalSpeedMod(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}



class NumberListContext extends antlr4.ParserRuleContext {

    constructor(parser, parent, invokingState) {
        if(parent===undefined) {
            parent = null;
        }
        if(invokingState===undefined || invokingState===null) {
            invokingState = -1;
        }
        super(parent, invokingState);
        this.parser = parser;
        this.ruleIndex = RasterklangDslParser.RULE_numberList;
    }

	NUMBER = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.NUMBER);
	    } else {
	        return this.getToken(RasterklangDslParser.NUMBER, i);
	    }
	};


	COMMA = function(i) {
		if(i===undefined) {
			i = null;
		}
	    if(i===null) {
	        return this.getTokens(RasterklangDslParser.COMMA);
	    } else {
	        return this.getToken(RasterklangDslParser.COMMA, i);
	    }
	};


	enterRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.enterNumberList(this);
		}
	}

	exitRule(listener) {
	    if(listener instanceof RasterklangDslListener ) {
	        listener.exitNumberList(this);
		}
	}

	accept(visitor) {
	    if ( visitor instanceof RasterklangDslVisitor ) {
	        return visitor.visitNumberList(this);
	    } else {
	        return visitor.visitChildren(this);
	    }
	}


}




RasterklangDslParser.ProgramContext = ProgramContext; 
RasterklangDslParser.StatementContext = StatementContext; 
RasterklangDslParser.UnknownStmtContext = UnknownStmtContext; 
RasterklangDslParser.UnknownTailContext = UnknownTailContext; 
RasterklangDslParser.AnyIdentContext = AnyIdentContext; 
RasterklangDslParser.StrContext = StrContext; 
RasterklangDslParser.TempoStmtContext = TempoStmtContext; 
RasterklangDslParser.TempoValueContext = TempoValueContext; 
RasterklangDslParser.TempoSettingContext = TempoSettingContext; 
RasterklangDslParser.TargetStmtContext = TargetStmtContext; 
RasterklangDslParser.ChipStmtContext = ChipStmtContext; 
RasterklangDslParser.TargetSettingContext = TargetSettingContext; 
RasterklangDslParser.PatternDeclContext = PatternDeclContext; 
RasterklangDslParser.PatternEachContext = PatternEachContext; 
RasterklangDslParser.ArrangementContext = ArrangementContext; 
RasterklangDslParser.ArrangementElementContext = ArrangementElementContext; 
RasterklangDslParser.TableDeclContext = TableDeclContext; 
RasterklangDslParser.InstrumentBlockContext = InstrumentBlockContext; 
RasterklangDslParser.InstrumentFieldContext = InstrumentFieldContext; 
RasterklangDslParser.TableVerbLineContext = TableVerbLineContext; 
RasterklangDslParser.TableRowContext = TableRowContext; 
RasterklangDslParser.TableTokenContext = TableTokenContext; 
RasterklangDslParser.AbsNoteContext = AbsNoteContext; 
RasterklangDslParser.RampGeneratorContext = RampGeneratorContext; 
RasterklangDslParser.VoiceStmtContext = VoiceStmtContext; 
RasterklangDslParser.VoiceBodyContext = VoiceBodyContext; 
RasterklangDslParser.VoiceShortTailContext = VoiceShortTailContext; 
RasterklangDslParser.InlineAttrContext = InlineAttrContext; 
RasterklangDslParser.ValueListContext = ValueListContext; 
RasterklangDslParser.VoiceFieldContext = VoiceFieldContext; 
RasterklangDslParser.VoiceFieldValueContext = VoiceFieldValueContext; 
RasterklangDslParser.PipelineStepContext = PipelineStepContext; 
RasterklangDslParser.TransformContext = TransformContext; 
RasterklangDslParser.TransformArgContext = TransformArgContext; 
RasterklangDslParser.TransformRefContext = TransformRefContext; 
RasterklangDslParser.FilterBlockContext = FilterBlockContext; 
RasterklangDslParser.FilterFieldContext = FilterFieldContext; 
RasterklangDslParser.FilterFieldValueContext = FilterFieldValueContext; 
RasterklangDslParser.SwitchCaseContext = SwitchCaseContext; 
RasterklangDslParser.SwitchCondContext = SwitchCondContext; 
RasterklangDslParser.SignalRangeOptContext = SignalRangeOptContext; 
RasterklangDslParser.SignalSpeedModContext = SignalSpeedModContext; 
RasterklangDslParser.NumberListContext = NumberListContext; 
