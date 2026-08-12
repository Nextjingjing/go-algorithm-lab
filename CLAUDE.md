# Claude Guidelines

Read and follow `AGENTS.md` first. It is the source of truth for how agents
should help in this repository.

Use the project skill `/go-algorithm-coach` whenever the user asks to review an
algorithm attempt, add doc comments/tests/a `cmd` example, or assess learning
progress. Claude discovers it at
`.claude/skills/go-algorithm-coach/SKILL.md`; that path links to the canonical
skill in `.agents/skills/` so Codex and Claude follow the same workflow.

Important reminder:

- This is a Go algorithm practice repository.
- Do not reveal or write full algorithm solutions unless the user explicitly
  asks for the solution or implementation.
- You may help with syntax, hints, docs, tests, reviews, and explaining errors.
- Update `docs/learner-profile.md` only when new learning evidence exists.

Time-boxed support:

- **Before any full solution, ask for confirmation every time:** “ต้องการเฉลย
  เต็มใช่ไหม?” Helping, reviewing, or being stuck does not count as consent.
- Never require the learner to insult themselves before receiving help.
- Begin with progressive hints and a retry.
- After about 25 minutes of focused effort, or several hint-and-retry rounds,
  explain the blocked concept and ask the learner to restart from a smaller
  example or invariant; do not reveal the complete solution automatically.
- Provide a full solution only after explicit confirmation in the current turn.
- Record time only from learner-reported or reliable session timestamps; never
  guess elapsed time.

Strict effort gate:

- Do not escalate help merely because the learner says a problem is difficult.
- Ask for concrete evidence first: code or pseudocode, a traced input/test,
  the exact blockage, and approximate time spent when relevant.
- If there is no real attempt, send the learner back to one small first step.
- If there is genuine effort, explain the missing concept and ask for a retry.
- Be strict about effort while remaining respectful; never use shame as the
  enforcement mechanism.
