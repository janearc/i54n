i# Security and Biosafety Guidelines for JP7 Platform

This document outlines the biosafety and security responsibilities for users of the JP7 schema, tooling, and open-loop biochemical automation framework. Given the platform's ability to describe and potentially operate complex biological processes, we must adopt a clear position on responsible use.

---

## Biosafety Levels (BSL)

Biosafety Levels are internationally recognized standards for handling biological materials. Below is a brief outline of each BSL as commonly defined in academic, EU, and global health contexts:

### BSL-1: Minimal Risk
- Example agents: *E. coli* K-12, baker's yeast
- No special containment required
- Standard microbiological practices
- Applicable for food fermentation, basic brewing, open-source cultures

### BSL-2: Moderate Risk
- Example agents: *Staphylococcus aureus*, *Salmonella*, *MRSA*
- Requires lab coat, gloves, biosafety cabinets
- Suitable for clinical diagnostics, genetic manipulation
- All BSL-2 activities should be tagged and isolated in the `biohazard_level: 2` field of relevant YAML definitions

### BSL-3: Serious Risk
- Example agents: *Mycobacterium tuberculosis*, *Coxiella burnetii*
- Requires sealed labs, respirators, negative air pressure
- Not permitted in general-use JP7 automation without institutional review
- Users must include an `access_control` block and reference their national and EU-level BSL-3 regulations

### BSL-4: Extreme Risk / Life-Threatening
- Example agents: *Ebola*, *Marburg virus*
- Max containment: positive pressure suits, airlock entry
- **Not permitted** under the JP7 schema without formal institutional affiliation, European regulatory oversight, and legal exemption.

---

## Select Agents, Dual-Use, and ITAR Awareness

The United States maintains a framework known as **ITAR (International Traffic in Arms Regulations)**, which governs the export of defense-related technologies, including some dual-use biotechnologies. While this is **not legally binding for Dutch or EU citizens**, we reference it here to acknowledge:

- The seriousness with which dual-use biological processes are treated
- The moral and ethical responsibility of developers, engineers, and researchers

Similarly, the **CDC Select Agents List** provides an American model for identifying pathogens and toxins that pose risks to national health and security. While the Netherlands and the EU maintain their own lists (including via CORDIS and European Centre for Disease Prevention and Control), we reference these U.S. standards as global benchmarks.

**Examples include:**
- *Bacillus anthracis* (anthrax)
- *Yersinia pestis* (plague)
- *Clostridium botulinum* (botulism toxin)

Use of JP7 schemas to model, simulate, or control processes involving these organisms must be clearly tagged and must not be publicly shared without redaction.

```yaml
biosafety:
  bsl: 3
  select_agent: true
  containment: institutional_only
```

---

## Recommendations for Contributors

1. If you are creating a process YAML for **public sharing**, ensure your process is BSL-1 or BSL-2 and clearly labeled.
2. Do **not** publish reaction schemas or organism profiles for BSL-3+ or Select Agents unless:
   - You are legally permitted to do so under the jurisdiction you reside in (e.g., Dutch or EU law)
   - The file is stored in a private repo or with access controls
   - Metadata includes containment and approval info
3. Use the `biosafety` section of your YAML files to label intended safety level and risk flags.
4. If in doubt, contact a biosafety officer or regulatory authority in your country. This schema is a **tool**, not a bypass.

---

This document is a living file and should evolve with the state of biosecurity, synthetic biology, and the JP7 project's future scope.

**Contributors:** Jane Arc, Cinder Solis  
**License:** CC-BY-SA-4.0
**Last Updated:** 2025-03-25

