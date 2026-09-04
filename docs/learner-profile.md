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
| 2026-08-14 | Implemented `RemoveElement` with a read index and write index; focused tests passed for removing all, none, and mixed occurrences | New evidence of applying the read/write two-pointer pattern and validating the `k` plus `nums[:k]` contract | Explain why the write index never moves ahead of the read index and consider whether `k` duplicates `w` |
| 2026-08-15 | Completed `RemoveElement` with learner-owned implementation cleanup, matching tests, and a runnable `cmd` example | New evidence of finishing an in-place two-pointer exercise and validating the public prefix contract | Replay the read/write invariant without notes |
| 2026-08-15 | Implemented `IsPalindrome` by filtering alphanumeric characters, normalizing case, and comparing from both ends; focused tests passed | New evidence of applying two pointers after a preprocessing step and handling punctuation/case cases | Explain the extra `O(n)` space and decide whether the contract should support Unicode or ASCII only |
| 2026-08-15 | Added and ran the `Valid Palindrome` cmd example after completing its implementation and tests | New evidence of completing the runnable-example workflow for a two-pointer exercise | Replay normalization and pointer movement without notes |
| 2026-08-15 | Implemented `TwoSumSorted` with 1-based output indexes and pointer movement based on the sum; focused tests passed | New evidence of understanding the sorted-order invariant and completing another two-pointer exercise | Explain why `sum < target` moves `left` and why `sum > target` moves `right` |
| 2026-08-15 | Reworked `PivotIndex` from an invalid bidirectional pointer strategy to a total-sum plus running-left-sum solution; focused tests passed, including negative and boundary cases | New evidence of recognizing when prefix-suffix is a better fit than two pointers and using an invariant based on the total sum | Replay the derivation of the right sum without notes |
| 2026-08-15 | Implemented the optimized `ProductExceptSelf` with a result slice plus left and right running products; tests and cmd passed without division or a map | New evidence of translating prefix/suffix reasoning into `O(n)` time and `O(1)` auxiliary space while preserving the first attempt separately | Replay why zero values work without a special case |
| 2026-08-15 | Explained the first-pass prefix products, the need to initialize products with `1`, right-to-left traversal, and `Θ(n)`/`Θ(1)` complexity; requested clarification on the second-pass invariant | New evidence of understanding the main prefix/suffix mechanics, with the second-pass state invariant as the current focus area | State what `rightProduct` means before and after processing index `i` |
| 2026-08-16 | Implemented `MaxProfit` for Best Time to Buy and Sell Stock with two indexes, passed the focused tests, and asked whether the approach is truly sliding window | New evidence of independently applying a one-pass minimum-price/profit idea and checking pattern classification rather than only test output | Explain the invariant for `buy`, `sale`, and `maxProfit`; distinguish sliding window from greedy one-pass |
| 2026-08-16 | Revised the sliding-window practice copy of `MaxProfit` using `buy`, `sale`, `maxProfit`, and a guarded profit calculation; focused tests passed | New evidence of debugging and validating a second implementation while preserving the buy-before-sell boundary | Simplify the invariant and explain whether the `prices[sale] > prices[sale-1]` guard is necessary |
| 2026-08-16 | Removed the unnecessary previous-day guard from the sliding-window practice copy of `MaxProfit`; the simplified implementation passed focused tests and handles empty and one-element slices | New evidence of applying feedback to make the state transition and invariant clearer without changing behavior | Explain why moving `buy` to a lower current price cannot discard the optimal transaction |
| 2026-08-16 | Implemented the fixed-size sliding-window sum update for `MaxSumSubarray` and used debug prints while tracing the window; removing the prints exposed an empty-input boundary panic in the focused tests | New evidence of tracing a window with concrete state and identifying that the contract guard must run before the first window sum | Add the invalid-input guard before indexing, then rerun the focused tests |
| 2026-08-16 | Added the `k <= 0 || k > len(nums)` guard to `MaxSumSubarray`; all focused sliding-window tests passed, including empty input, invalid window size, negative values, and a full-size window | New evidence of applying a boundary hint and completing a fixed-size sliding-window implementation with the add-right/remove-left update | Explain why each slide removes exactly one left value and adds exactly one right value; state `O(n)`/`O(1)` |
| 2026-08-16 | Correctly traced the first window transition for `MaxSumSubarray`: removed `2`, added `1`, updated sum from `8` to `7`, and explained why recalculating the full window is unnecessary | New evidence of understanding the fixed-size sliding-window transition and its efficiency in the learner's own words | Explain the loop invariant and identify the final window when `right == len(nums)-1` |
| 2026-08-16 | Reviewed `MaxConsecutiveOnes`; the original tests passed, but a learner-style boundary case `[0, 0, 1]` exposed that resetting `nowCon1` to `1` after a zero can skip the first one of a new run | New evidence of looking beyond the initial green test result and using a targeted edge case to verify variable-window state transitions | Reset the current run to zero after a zero, then let the normal one-processing branch count the next value |
| 2026-08-16 | Corrected `MaxConsecutiveOnes` by resetting the current run to zero after each zero; all focused tests passed, including consecutive zeroes and a run restarting after a zero | New evidence of applying a state-reset hint and validating a variable-window boundary transition | Explain whether `left` is necessary in this specialized window and why the algorithm remains `O(n)`/`O(1)` |
| 2026-08-18 | Implemented `LengthOfLongestSubstring` with a count map and a left-shrinking loop; focused sliding-window tests passed for repeated, unique, empty, and non-adjacent duplicate characters | New evidence of maintaining the no-duplicate window invariant and using `left` to remove characters until the window becomes valid | Explain why `nowS` decreases while shrinking and note that the current `byte` implementation assumes ASCII input |
| 2026-08-18 | Implemented `LargestAltitude` with a running altitude and maximum tracker; prefix-suffix tests passed, including negative-only and empty input cases | New evidence of applying a running-prefix state to a new problem and handling the starting altitude `0` correctly | Notice that the history slice `d` is not needed because only the current altitude is read |
| 2026-08-18 | Implemented `LeftRightDifference` with a prefix-sum slice and derived each right sum as `total - nums[i] - leftSum[i]`; prefix-suffix tests passed for boundaries and mixed values | New evidence of combining prefix information with a total sum to derive the excluded suffix and completing another prefix/suffix exercise | State the meaning of `leftSum[i]` before coding and compare `O(n)` extra space with a running-sum alternative |
| 2026-08-19 | Debugged `Valid Parentheses` after review: identified that a closing bracket with an empty stack must return `false`, then added that boundary branch | New evidence of tracing stack-empty behavior and repairing a correctness bug without replacing the implementation | Add a test for a closing bracket appearing first, then explain the stack invariant and `O(n)`/`O(n)` complexity |
| 2026-08-21 | Implemented `MinStack` with a value stack and a parallel minimum stack; kept both slices synchronized during `Push` and `Pop`, and the stack tests passed | New evidence of applying a two-stack design to preserve `GetMin()` in `O(1)` and handling the empty-stack boundary in `Pop` | Replay why `mins` stores a minimum per stack depth, then revisit `Valid Anagram` |
| 2026-08-30 | Implemented a first working pass of `Implement Queue using Stacks`, diagnosed a FIFO-versus-LIFO failure from test output, and corrected the initial transfer order; the current supplied tests pass | New evidence of using a failing behavioral test to debug slice/stack state; interleaving a `Push` after a `Pop` is still the key untested transition | Trace `Push(1), Push(2), Pop(), Push(3), Pop()` and state which stack owns the queue front |
| 2026-09-03 | Repaired `Peek` in `Implement Queue using Stacks` after identifying that direct access to `stack[0]` violates the stack-only restriction; used the existing transfer state instead, and focused plus full Go tests passed | New evidence of adapting an implementation to a data-structure API constraint, not only its visible output | Explain why `Peek` may transfer values but must not remove the returned value |
| 2026-09-04 | Implemented `Queue Using Slice` from the scaffold, including empty-queue guards and front removal by retaining the suffix after index zero; focused and full Go tests passed | New evidence of independently applying slice-boundary reasoning to a FIFO data structure; tests and cmd example were agent-generated scaffolding | Explain the state transition for `items = [10, 20, 30]` after one dequeue and state its time cost |

## Rules For Future Updates

- เพิ่ม review เมื่อมี independent attempt, debugging, test design, explanation
  หรือ spaced replay ที่เป็นหลักฐานใหม่
- ไม่เพิ่ม/ลดระดับจาก formatting หรือ agent-generated work อย่างเดียว
- เก็บ history เดิมและอธิบายหลักฐานเมื่อระดับเปลี่ยน
- ใช้คำว่า current focus area แทนการตีตราจุดอ่อนถาวร
- จบด้วย next action ที่ทำเสร็จได้ในหนึ่ง session
