# Learner Profile

ไฟล์นี้เป็นสมุดบันทึกพัฒนาการ ไม่ใช่ใบเกรดหรือการตัดสินความฉลาด คะแนนหมายถึง
ระดับหลักฐานที่เห็นใน repository และบทสนทนาเท่านั้น และเปลี่ยนได้เสมอ

## Evidence Scale

| Level | Meaning |
| --- | --- |
| 0 | ยังไม่มีหลักฐานพอประเมิน |
| 1 | ทำได้เมื่อมีตัวอย่างหรือการนำค่อนข้างมาก |
| 2 | ทำได้ด้วย hint/review และแก้ต่อเองได้ |
| 3 | ทำได้เองในโจทย์ระดับปัจจุบันและอธิบายเหตุผลได้ |
| 4 | replay ได้หลังเว้นระยะ เปรียบเทียบทางเลือก และสอนกลับได้ |

## Snapshot — 2026-08-12

| Skill | Evidence level | Evidence |
| --- | ---: | --- |
| Independent implementation | 2 | เริ่มเขียน Linear Search, Two Sum, Merge Sort และโจทย์พื้นฐานหลายข้อก่อนขอ review; algorithm หลักมักมาจากความพยายามของผู้เรียน |
| Boundary reasoning | 2 | เริ่มระบุได้เองว่าปัญหาหลักเกี่ยวกับ `<`/`<=`; ยังถามซ้ำใน inclusive ranges, `mid`, และ `i+1` |
| Go slices, arrays, pointers | 2 | เข้าใจแล้วว่า slice แก้ backing array ได้โดยไม่ต้องใช้ `*[]int`; ยังต้องฝึกเลือก API ให้ idiomatic โดยไม่ใช้ pointer-to-slice เกินจำเป็น |
| Debugging | 2 | นำ compiler error และบรรทัดที่สงสัยมาถามได้ตรงจุด แล้วแก้ต่อเป็นรอบสั้น ๆ |
| Test design | 1 | อ่านผล `go test` และเข้าใจเรื่อง cache/`[no test files]` มากขึ้น แต่ tests ส่วนใหญ่ให้ agent สร้าง จึงยังไม่มีหลักฐานว่าออกแบบ edge cases เองได้สม่ำเสมอ |
| Explanation and complexity | 1 | มีหลักฐานการคุยเรื่อง boundaries และ mutation แต่ยังไม่เห็นการอธิบาย invariant และ Big-O ด้วยตัวเองอย่างสม่ำเสมอ |
| Retention after spacing | 0 | ยังไม่มี no-notes replay ที่บันทึกหลังเว้น 7 หรือ 21 วัน |

## What Is Going Well

- ลงมือเขียนก่อนถามในโจทย์ algorithm หลายข้อ ไม่ได้ขอ full solution เป็นนิสัย
- คำถามสั้นแต่เจาะจง เช่น compiler error, pointer semantics และ loop boundary
- ยอมย้อนกลับมาถาม concept เดิมจนเข้าใจ ไม่แกล้งผ่าน
- ทำรอบ feedback เร็ว: เขียน → ขอ review → แก้ → ขอเช็กอีกครั้ง
- ความช่วยเหลือที่ขอบ่อยส่วนใหญ่เป็น syntax, concept, review, docs, tests และ
  `cmd/` ซึ่งเป็น scaffolding ที่มีประโยชน์ ไม่ใช่สัญญาณว่าเรียนไม่ได้

## Current Focus Areas

1. **Boundary model** — เขียนช่วง index บนกระดาษก่อน loop ทุกครั้ง เช่น
   `[left..right]` หรือ `[left:right)` แล้วค่อยเลือก `<`/`<=`.
2. **Contract before code** — ตัดสินใจให้ชัดว่า empty input คืนอะไร, panic หรือ
   error ก่อน implementation.
3. **Test ownership** — ก่อนขอให้ agent เขียน test ให้เสนอ case เองอย่างน้อย 3
   case: normal, boundary และ failure/not-found.
4. **Go API judgment** — ใช้ `[]T` เมื่อแก้ element; ใช้ `*[]T` เฉพาะเมื่อต้อง
   เปลี่ยน slice header ของ caller และการ return slice ไม่เหมาะกว่า.
5. **Explain the invariant** — ฝึกประโยค “ก่อนเริ่มรอบนี้ ส่วนไหนของข้อมูลถูกต้อง
   แล้ว” เพื่อเปลี่ยนจากจำ code เป็นเข้าใจ pattern.

## Current Code Evidence To Revisit

- `BinarySearch` ยังไม่มี base case สำหรับช่วงที่หมด (`left > right`) และ empty
  slice จึงเป็น debugging checkpoint ที่ดี; tests ปัจจุบันยังไม่ครอบคลุม empty,
  one-element not found และ target ที่มากกว่าค่าสุดท้าย
- `FindMax` ใช้ `*[]int` ทั้งที่อ่าน element ผ่าน `[]int` ได้ และ contract กำหนด
  ว่าห้าม empty; เหมาะสำหรับฝึกออกแบบ function signature
- `TwoSum` ทำ brute force ได้แล้ว แต่ยังไม่มี Go doc comment; รอบ hash map ควร
  เริ่มจากอธิบาย trade-off ก่อนลงมือ
- `Valid Anagram` ทำ hash-map version แล้ว โดยนับจำนวนตัวอักษรและเพิ่ม tests กับ
  `cmd/` example; ควร replay จากความจำและพิจารณาเงื่อนไข Unicode ต่อไป
- tests ที่มีเป็น table-driven และครอบคลุมพื้นฐานค่อนข้างดี แต่เป็นงานที่ agent
  ช่วยสร้าง จึงควรเปลี่ยนบทบาทครั้งถัดไป: ผู้เรียนออกแบบ cases, agent ตรวจ

## Help Pattern

ความถี่ในการขอความช่วยเหลือค่อนข้างสูง แต่รูปแบบปัจจุบัน productive:

- บ่อย: syntax, pointer/slice model, boundaries, review
- บ่อย: ให้ช่วยสร้าง doc comments, tests และ `cmd/`
- น้อย: ขอให้เขียน algorithm solution เต็มให้ตั้งแต่ต้น

เป้าหมายไม่ใช่ลดจำนวนคำถามทันที แต่เปลี่ยนคำถามจาก “เขียน test ให้” เป็น
“ฉันเลือก cases เหล่านี้ เพราะอะไรขาดไหม” และจาก “ใช้ `<` หรือ `<=`” เป็น
“ฉันเลือก `<` เพราะช่วงของฉันเป็นแบบนี้ ถูกไหม”

## Sustainable Help Rule

ผู้เรียนเลือกใช้ time-box แทนการบังคับตัวเองด้วยคำพูดดูถูก: เริ่มจากคำใบ้เล็ก ๆ
แล้วลองใหม่ หากใช้ความพยายามต่อเนื่องประมาณ 25 นาทีหรือผ่านหลายรอบของ
hint-and-retry แล้วยังติด ให้เปลี่ยนเป็นการอธิบายจุดติดและเริ่มโจทย์ใหม่จากเคสที่
เล็กลงหรือเขียน invariant ก่อน โดยต้องถามยืนยันอีกครั้งก่อนเฉลยเต็มทุกครั้ง

> **กติกาเฉลยเต็ม:** ต้องถามผู้เรียนก่อนทุกครั้งว่า “ต้องการเฉลยเต็มใช่ไหม?”
> การขอความช่วยเหลือ การขอ review หรือการติดนานไม่ถือเป็นการอนุญาต

การบันทึกเวลาให้ใช้เฉพาะเวลาที่ผู้เรียนรายงานหรือมี timestamp ที่เชื่อถือได้
เท่านั้น ไม่เดาระยะเวลา และไม่ใช้ self-insult เป็นเงื่อนไขการได้รับความช่วยเหลือ

## Strict Effort Gate

ความยากอย่างเดียวไม่ใช่เหตุผลพอที่จะขอ hint ระดับสูงขึ้น ก่อนช่วยต่อควรมีหลักฐาน
อย่างน้อย: โค้ดหรือ pseudocode ที่ลอง, input/test ที่ trace แล้วหนึ่งเคส, จุดที่
เหตุผลขาด และเวลาคร่าว ๆ ที่ใช้เมื่อเกี่ยวข้อง ถ้ายังไม่ได้พยายาม ให้กลับไปทำ
ขั้นเล็ก ๆ ก่อน ถ้าพยายามจริงแต่ติด จึงค่อยสอน concept ที่ขาดและให้ลองใหม่

กติกานี้เข้มกับ “ความพยายาม” แต่ไม่ดูถูก “ตัวคน” เพื่อแยกการติดเพราะโจทย์ยาก
ออกจากการข้ามขั้นเพราะไม่อยากคิด

## Next Small Action

Replay `Valid Anagram` จากความจำ โดยเขียน invariant หนึ่งประโยคว่า map เก็บอะไร
และจด test cases เอง 3 case ก่อนให้ agent review

## Review History

| Date | Evidence | Change | Next focus |
| --- | --- | --- | --- |
| 2026-08-12 | Repository algorithms, tests, and coaching conversation to date | Created baseline; no retention score yet | Own test-case design and boundary contracts |
| 2026-08-12 | Reworked `MoveZeroes` from repeated shifting to a one-pass two-pointer implementation after guided hints; matching tests pass | New evidence of applying a named pattern and preserving in-place behavior | Explain the two pointer invariant and `O(n)` complexity without notes |
| 2026-08-12 | Traced `[0, 1, 0, 3, 12]` through each scan and correctly explained that `j` scans while `i` marks the next non-zero position | Stronger evidence of boundary/state reasoning for two pointers | Replay `MoveZeroes` tomorrow without notes |
| 2026-08-13 | Added a self-designed `SumArray` case for a slice containing only zeros, with the expected sum `0` | First recorded evidence of learner-owned table-test case design | Run the focused test and add one boundary-oriented case |
| 2026-08-13 | Correctly identified `SumArray` time complexity as `Θ(n)`; initially explained empty input using the all-zero case and received a distinction between the two contracts | New evidence of beginning complexity explanation and boundary clarification | Explain why the auxiliary space is constant |
| 2026-08-13 | Implemented `ContainsDuplicate` with a map and early return when a count exceeds one; the implementation handles empty input and repeated values | New evidence of independently applying the hash-map pattern after learning basic Go map syntax | Explain the map invariant and compare `map[int]int` with a presence map |
| 2026-08-13 | Implemented `Valid Anagram` with two character-count maps, corrected the `byte` versus `string` key mismatch, and discussed `O(n)` time and `O(k)` space | New evidence of applying counting-map logic and debugging a Go map type error with guided help; tests and cmd example were agent-generated | Replay the exercise and explain the character-count invariant without notes |
| 2026-08-14 | Designed additional `Valid Anagram` cases for case sensitivity and spaces; identified that spaces are counted as characters and that unequal lengths must return `false` | New evidence of learner-owned test-case design and contract reasoning; two initial space expectations were incorrect and exposed by focused tests | Correct the space cases and explain whether the contract ignores whitespace or treats it as input |
| 2026-08-14 | Added learner-owned `BinarySearch` cases for an empty slice and targets outside the sorted range; the empty-slice case exposed an index-out-of-range panic | New evidence of boundary-oriented test design and using a failing test to locate an exhausted-range bug | State the recursive base case for `left > right` before editing the implementation |
| 2026-08-14 | Added the exhausted-range base case to `BinarySearch`; the focused test command passed after the empty-slice failure | New evidence of applying a boundary invariant to repair a recursive algorithm and validating the fix independently | Add or explain right-edge and one-element cases, then run the package-level test |
| 2026-08-14 | Added `BinarySearch` cases for a one-element slice and the right edge; package-level tests passed | New evidence of completing boundary-oriented test coverage and validating the repaired implementation | Add the one-element-not-found case and avoid tests that require a specific duplicate index unless the contract guarantees it |
| 2026-08-14 | Implemented `MinValue` by initializing from the first element, scanning the remaining values, and handling negative values; focused tests passed | New evidence of correct initialization under an explicit non-empty precondition and independent implementation of a linear scan | Explain `O(n)` time and `O(1)` auxiliary space; improve the empty-input panic message if keeping that guard |

## Rules For Future Updates

- เพิ่ม review เมื่อมี independent attempt, debugging, test design, explanation
  หรือ spaced replay ที่เป็นหลักฐานใหม่
- ไม่เพิ่ม/ลดระดับจาก formatting หรือ agent-generated work อย่างเดียว
- เก็บ history เดิมและอธิบายหลักฐานเมื่อระดับเปลี่ยน
- ใช้คำว่า current focus area แทนการตีตราจุดอ่อนถาวร
- จบด้วย next action ที่ทำเสร็จได้ในหนึ่ง session
